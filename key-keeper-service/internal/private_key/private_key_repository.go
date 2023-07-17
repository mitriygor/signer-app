package private_key

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"key-keeper-service/config"
	ah "key-keeper-service/pkg/args_helper"
	"log"
	"strconv"
	"strings"
)

type Repository interface {
	GetKeys(args KeyPayload) ([]*PrivateKey, error)
	IncrCount(countName string)
	SetCount(count int, countName string)
	GetCount(countName string) int
}

type PrivateKeyRepository struct {
	DB          *sql.DB
	redisClient *redis.Client
}

func NewPrivateKeyRepository(db *sql.DB, redisClient *redis.Client) Repository {
	return &PrivateKeyRepository{
		DB:          db,
		redisClient: redisClient,
	}
}

func (pk *PrivateKeyRepository) GetKeys(args KeyPayload) ([]*PrivateKey, error) {
	sqlQueryBuilder := strings.Builder{}
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	sqlQueryBuilder.WriteString(`SELECT id, title, secret FROM private_key WHERE 1 = 1`)

	if args.KeyLimit != 0 {
		ah.AddArgToQuery(&sqlQueryBuilder, "LIMIT", strconv.Itoa(args.KeyLimit))
	}

	rows, err := pk.DB.QueryContext(ctx, sqlQueryBuilder.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	privateKeys := make([]*PrivateKey, 0)

	for rows.Next() {
		privateKey := PrivateKey{}
		err := rows.Scan(&privateKey.ID, &privateKey.Title, &privateKey.Secret)
		if err != nil {
			pk.IncrCount(config.ErrorCount)
			fmt.Printf("\nBroker :: DB :: GetKeys :: ERROR:%v\n", err.Error())
		} else {
			pk.IncrCount(config.ReqCount)
			privateKeys = append(privateKeys, &privateKey)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return privateKeys, nil
}

func (pk *PrivateKeyRepository) IncrCount(countName string) {
	pk.redisClient.Incr(context.TODO(), countName)
}

func (pk *PrivateKeyRepository) SetCount(count int, countName string) {
	pk.redisClient.Set(context.TODO(), countName, count, 0)
}

func (pk *PrivateKeyRepository) GetCount(countName string) int {
	count, err := pk.redisClient.Get(context.TODO(), countName).Int()
	if err != nil {
		fmt.Printf("\nBroker :: Redis :: GetCount :: ERROR:%v\n", err.Error())
		return -1
	}

	return count
}
