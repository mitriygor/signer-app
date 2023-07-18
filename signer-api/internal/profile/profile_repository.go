package profile

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"os"
	"signer-api/config"
	"signer-api/internal/broker"
	"signer-api/internal/private_key"
	ah "signer-api/pkg/args_helper"
	"signer-api/pkg/sign_helper"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	privateKeys []*private_key.PrivateKey
	batchStamp  string
)

type Repository interface {
	GetAll(args Args) ([]*Profile, error)
	SignAllWithParams(signPayload SignPayload) error
	SignAll() error
	SignBatch(profileBatchChan <-chan ProfileBatch, wg *sync.WaitGroup)
	SignProfile(profile Profile, privateKey *private_key.PrivateKey, signature string, wgInner *sync.WaitGroup) error
	IncrCount(countName string)
	SetCount(count int, countName string)
	GetCount(countName string) int
}

type profileRepository struct {
	DB                   *sql.DB
	DBW                  *sql.DB
	PrivateKeyRepository private_key.Repository
	BrokerService        broker.Service
	redisClient          *redis.Client
	mu                   sync.Mutex
}

func NewProfileRepository(db *sql.DB, pkr private_key.Repository, dbw *sql.DB, brokerService broker.Service, redisClient *redis.Client) Repository {
	return &profileRepository{
		DB:                   db,
		DBW:                  dbw,
		PrivateKeyRepository: pkr,
		BrokerService:        brokerService,
		redisClient:          redisClient,
	}
}

type ProfileBatch struct {
	privateKey *private_key.PrivateKey
	profiles   []Profile
}

var file, _ = os.Create("insert.psql")

var sem = make(chan int, 10)

func (pr *profileRepository) GetAll(args Args) ([]*Profile, error) {
	sqlQueryBuilder := strings.Builder{}
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	sqlQueryBuilder.WriteString(`SELECT id, first_name, last_name FROM profile WHERE 1 = 1`)

	if args.FirstName != "" {
		ah.AddArgToQuery(&sqlQueryBuilder, "AND first_name LIKE", args.FirstName)
	}

	if args.LastName != "" {
		ah.AddArgToQuery(&sqlQueryBuilder, "AND last_name LIKE", args.LastName)
	}

	if args.Limit != 0 {
		ah.AddArgToQuery(&sqlQueryBuilder, "LIMIT", strconv.Itoa(args.Limit))
	} else {
		ah.AddArgToQuery(&sqlQueryBuilder, "LIMIT", "10")
	}

	rows, err := pr.DB.QueryContext(ctx, sqlQueryBuilder.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	profiles := make([]*Profile, 0)

	for rows.Next() {
		profile := Profile{}
		err := rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName)
		if err != nil {
			log.Fatal(err)
		}
		profiles = append(profiles, &profile)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return profiles, nil
}

func (pr *profileRepository) SignAll() error {

	var err error

	privateKeys, err = pr.getPrivateKeys()
	if err != nil {

		return err
	}

	numsWorkers := pr.getNumsWorkers()

	totalProfiles, err := pr.getTotalProfiles()
	if err != nil {

		return err
	}

	batchStamp = sign_helper.GetStamp()

	batchSize := pr.getBatchSize()

	numBatches := (totalProfiles + batchSize - 1) / batchSize

	var wg sync.WaitGroup

	profileBatchChan := make(chan ProfileBatch, numsWorkers)

	for i := 0; i < numsWorkers; i++ {
		wg.Add(1)
		go pr.SignBatch(profileBatchChan, &wg)
	}

	for i := 0; i < numBatches; i++ {

		offset := i * batchSize
		query := fmt.Sprintf("SELECT * FROM profile LIMIT %d OFFSET %d", batchSize, offset)

		rows, err := pr.DB.QueryContext(context.Background(), query)
		if err != nil {

			log.Fatal(err)
		}
		defer rows.Close()

		profiles := make([]Profile, 0)

		for rows.Next() {
			var profile Profile
			var profileSignature, profileStamp sql.NullString
			var profilePrivateKeyID sql.NullInt64
			var updatedAt sql.NullTime
			err := rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profileSignature, &profileStamp, &profilePrivateKeyID, &updatedAt)
			if err != nil {

				log.Fatal(err)
			}

			if profileSignature.Valid {
				profile.Signature = string(profileSignature.String)
			}
			if profileStamp.Valid {
				profile.Stamp = string(profileStamp.String)
			}
			if profilePrivateKeyID.Valid {
				profile.PrivateKeyID = int(profilePrivateKeyID.Int64)
			}
			if updatedAt.Valid {
				profile.UpdatedAt = updatedAt.Time
			}

			profiles = append(profiles, profile)
		}

		if err := rows.Err(); err != nil {

			log.Fatal(err)
		}

		var privateKey *private_key.PrivateKey
		privateKey, privateKeys = pr.selectPrivateKey()

		profileBatch := ProfileBatch{
			privateKey: privateKey,
			profiles:   profiles,
		}

		profileBatchChan <- profileBatch
	}

	close(profileBatchChan)
	wg.Wait()

	return nil
}

func (pr *profileRepository) SignAllWithParams(signPayload SignPayload) error {

	privateKeys = signPayload.Keys
	numsWorkers := signPayload.WorkersAmount
	totalProfiles := signPayload.RecordsAmount
	batchSize := signPayload.BatchSize
	numBatches := (totalProfiles + batchSize - 1) / batchSize

	batchStamp = sign_helper.GetStamp()

	var wg sync.WaitGroup

	profileBatchChan := make(chan ProfileBatch, numsWorkers)

	for i := 0; i < numsWorkers; i++ {
		wg.Add(1)
		go pr.SignBatch(profileBatchChan, &wg)
	}

	for i := 0; i < numBatches; i++ {

		offset := i * batchSize
		query := fmt.Sprintf("SELECT * FROM profile LIMIT %d OFFSET %d", batchSize, offset)

		rows, err := pr.DB.QueryContext(context.Background(), query)
		if err != nil {

			log.Fatal(err)
		}
		defer rows.Close()

		profiles := make([]Profile, 0)

		for rows.Next() {
			var profile Profile
			var profileSignature, profileStamp sql.NullString
			var profilePrivateKeyID sql.NullInt64
			var updatedAt sql.NullTime
			err := rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profileSignature, &profileStamp, &profilePrivateKeyID, &updatedAt)
			if err != nil {

				log.Fatal(err)
			}

			if profileSignature.Valid {
				profile.Signature = string(profileSignature.String)
			}
			if profileStamp.Valid {
				profile.Stamp = string(profileStamp.String)
			}
			if profilePrivateKeyID.Valid {
				profile.PrivateKeyID = int(profilePrivateKeyID.Int64)
			}
			if updatedAt.Valid {
				profile.UpdatedAt = updatedAt.Time
			}

			profiles = append(profiles, profile)
		}

		if err := rows.Err(); err != nil {

			log.Fatal(err)
		}

		var privateKey *private_key.PrivateKey
		privateKey, privateKeys = pr.selectPrivateKey()

		profileBatch := ProfileBatch{
			privateKey: privateKey,
			profiles:   profiles,
		}

		profileBatchChan <- profileBatch
	}

	close(profileBatchChan)
	wg.Wait()

	return nil
}

func (pr *profileRepository) SignBatch(profileBatchChan <-chan ProfileBatch, wg *sync.WaitGroup) {

	var wgInner sync.WaitGroup

	defer wg.Done()

	for profileBatch := range profileBatchChan {

		for _, profile := range profileBatch.profiles {
			wgInner.Add(1)

			privateKey := profileBatch.privateKey
			go func(p Profile) {

				if batchStamp != p.Stamp {

					signature := sign_helper.Encode(strconv.Itoa(p.ID), profileBatch.privateKey.Secret)

					err := pr.SignProfile(p, privateKey, signature, &wgInner)

					if err != nil {
						pr.logEntity("ERROR", signature, p.ID, privateKey.ID, "")
					} else {
						pr.logEntity("INFO", signature, p.ID, privateKey.ID, "")
					}
				}
			}(profile)
		}
	}

	wgInner.Wait()
}

func (pr *profileRepository) SignProfile(profile Profile, privateKey *private_key.PrivateKey, signature string, wgInner *sync.WaitGroup) error {

	defer wgInner.Done()

	privateKey.Mutex.Lock()
	profile.Mutex.Lock()

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("INSERT INTO profile (id, first_name, last_name, signature, stamp, private_key_id, updated_at) ")
	queryBuilder.WriteString("VALUES ( ")
	queryBuilder.WriteString(strconv.Itoa(profile.ID) + ", ")
	queryBuilder.WriteString("'" + profile.FirstName + "', ")
	queryBuilder.WriteString("'" + profile.LastName + "', ")
	queryBuilder.WriteString("'" + signature + "', ")
	queryBuilder.WriteString("'" + batchStamp + "', ")
	queryBuilder.WriteString(strconv.Itoa(privateKey.ID) + ", ")
	queryBuilder.WriteString("'" + time.Now().Format("2006-01-02 15:04:05.999999999") + "' ) ")
	queryBuilder.WriteString("ON CONFLICT (id) DO UPDATE ")
	queryBuilder.WriteString("SET ")
	queryBuilder.WriteString("first_name = EXCLUDED.first_name, ")
	queryBuilder.WriteString("last_name = EXCLUDED.last_name, ")
	queryBuilder.WriteString("signature = EXCLUDED.signature, ")
	queryBuilder.WriteString("stamp = EXCLUDED.stamp, ")
	queryBuilder.WriteString("private_key_id = EXCLUDED.private_key_id, ")
	queryBuilder.WriteString("updated_at = EXCLUDED.updated_at ")
	queryBuilder.WriteString("WHERE COALESCE(profile.stamp, '') <> '" + batchStamp + "';")

	query := queryBuilder.String()

	sem <- 1
	result, err := pr.DBW.Exec(query)
	<-sem

	if err != nil {
		fmt.Printf("\nSignProfile :: error :: result: %v\n", err.Error())
		pr.IncrCount(config.ErrorCount)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("\nSignProfile :: error:: rowsAffected: %v\n", err.Error())
		pr.IncrCount(config.ErrorCount)
		return err
	}

	if rowsAffected <= 0 {
		fmt.Printf("\nSignProfile :: error: NO UPDATE\n")
		pr.IncrCount(config.ErrorCount)
		return errors.New("nothing to update")
	}

	defer profile.Mutex.Unlock()
	defer privateKey.Mutex.Unlock()

	pr.IncrCount(config.ReqCount)

	return nil
}

func writeToFile(str string) {
	file.WriteString(str)
	file.WriteString(";\n")
}

func (pr *profileRepository) getPrivateKeys() ([]*private_key.PrivateKey, error) {
	privateKeys, err := pr.PrivateKeyRepository.GetAll(private_key.Args{})

	if err != nil {
		return nil, err
	}

	return privateKeys, nil
}

func (pr *profileRepository) getNumsWorkers() int {
	numsWorkers := config.NumsWorkers
	return numsWorkers
}

func (pr *profileRepository) getTotalProfiles() (int, error) {
	count := config.NumsProfiles
	query := "SELECT COUNT(*) FROM profile"
	err := pr.DB.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (pr *profileRepository) getBatchSize() int {
	batchSize := config.BatchSize
	return batchSize
}

func (pr *profileRepository) selectPrivateKey() (*private_key.PrivateKey, []*private_key.PrivateKey) {

	key := privateKeys[0]

	privateKeys = append(privateKeys[1:], key)

	return key, privateKeys
}

func (pr *profileRepository) logEntity(logType string, signature string, profileID int, privateKeyID int, data string) {

	logPayload := broker.LogPayload{
		Name:      "log",
		Type:      logType,
		Stamp:     batchStamp,
		Signature: signature,
		ProfileID: profileID,
		KeyID:     privateKeyID,
		Data:      data,
	}

	payload := broker.RequestPayload{
		Action: "log",
		Log:    logPayload,
	}

	pr.BrokerService.HandleQueue(payload)
}

func (pr *profileRepository) IncrCount(countName string) {
	pr.redisClient.Incr(context.TODO(), countName)
}

func (pr *profileRepository) SetCount(count int, countName string) {
	pr.redisClient.Set(context.TODO(), countName, count, 0)
}

func (pr *profileRepository) GetCount(countName string) int {
	count, err := pr.redisClient.Get(context.TODO(), countName).Int()
	if err != nil {
		fmt.Printf("\nSigner :: Redis :: GetCount :: error:%v\n", err.Error())
		return -1
	}

	return count
}
