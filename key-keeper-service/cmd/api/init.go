package main

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"key-keeper-service/config"
	"key-keeper-service/internal/private_key"
	"net/http"
)

func initHandlers(conn *sql.DB, redisClient *redis.Client) (http.Handler, error) {

	privateKeyRepo := private_key.NewPrivateKeyRepository(conn, redisClient)
	privateKeyService := private_key.NewPrivateKeyService(privateKeyRepo)
	privateKeyHandler := &private_key.Handler{PrivateKeyService: privateKeyService}

	privateKeyRepo.SetCount(0, config.ReqCount)
	privateKeyRepo.SetCount(0, config.ErrorCount)

	return Routes(*privateKeyHandler), nil
}
