package log_entry

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"logger-service/config"
	"time"
)

type Repository interface {
	Insert(logEntry LogEntry) error
	All() ([]*LogEntry, error)
	GetOne(id string) (*LogEntry, error)
	Update(logEntry LogEntry) (*mongo.UpdateResult, error)
	IncrCount(countName string)
	SetCount(count int, countName string)
	GetCount(countName string) int
}

type LogEntryRepository struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func NewLogEntryRepository(mongoClient *mongo.Client, redisClient *redis.Client) Repository {
	return &LogEntryRepository{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

func (le *LogEntryRepository) Insert(logEntry LogEntry) error {
	collection := le.mongoClient.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      logEntry.Name,
		Data:      logEntry.Data,
		Type:      logEntry.Type,
		Stamp:     logEntry.Stamp,
		Signature: logEntry.Signature,
		ProfileID: logEntry.ProfileID,
		KeyID:     logEntry.KeyID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into logs:", err)
		le.IncrCount(config.ErrorCount)
		return err
	}
	le.IncrCount(config.ReqCount)

	return nil
}

func (le *LogEntryRepository) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := le.mongoClient.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Print("Error decoding log into slice:", err)
			return nil, err
		} else {
			logs = append(logs, &item)
		}
	}

	return logs, nil
}

func (le *LogEntryRepository) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := le.mongoClient.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (le *LogEntryRepository) Update(logEntry LogEntry) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := le.mongoClient.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(logEntry.ID)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", logEntry.Name},
				{"data", logEntry.Data},
				{"updated_at", time.Now()},
			}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (le *LogEntryRepository) IncrCount(countName string) {
	le.redisClient.Incr(context.TODO(), countName)
}

func (le *LogEntryRepository) SetCount(count int, countName string) {
	le.redisClient.Set(context.TODO(), countName, count, 0)
}

func (le *LogEntryRepository) GetCount(countName string) int {
	count, err := le.redisClient.Get(context.TODO(), countName).Int()
	if err != nil {
		fmt.Printf("\nLogger :: Redis :: GetCount :: error: %v\n", err.Error())
		return -1
	}
	return count
}
