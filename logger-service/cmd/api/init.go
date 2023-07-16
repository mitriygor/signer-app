package main

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"logger-service/config"
	"logger-service/internal/log_entry"
	"net/http"
)

func initHandlers(mongoClient *mongo.Client, redisClient *redis.Client) (http.Handler, error) {
	logEntryRepo := log_entry.NewLogEntryRepository(mongoClient, redisClient)
	logEntryService := log_entry.NewLogEntryService(logEntryRepo)
	logEntryHandler := &log_entry.Handler{LogEntryService: logEntryService}

	logEntryRepo.SetCount(0, config.ReqCount)
	logEntryRepo.SetCount(0, config.ErrorCount)

	return Routes(*logEntryHandler), nil
}
