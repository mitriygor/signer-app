package main

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"net/http"
	"signer-api/config"
	"signer-api/internal/broker"
	"signer-api/internal/private_key"
	"signer-api/internal/profile"
)

func initHandlers(conn *sql.DB, connWriter *sql.DB, redisClient *redis.Client) (http.Handler, error) {
	privateKeyRepo := private_key.NewPrivateKeyRepository(conn)
	privateKeyService := private_key.NewPrivateKeyService(privateKeyRepo)
	privateKeyHandler := &private_key.Handler{PrivateKeyService: privateKeyService}

	brokerService := broker.NewBrokerService()
	brokerHandler := &broker.Handler{BrokerService: brokerService}

	profileRepo := profile.NewProfileRepository(conn, privateKeyRepo, connWriter, brokerService, redisClient)
	profileService := profile.NewProfileService(profileRepo)
	profileHandler := &profile.Handler{ProfileService: profileService}

	profileRepo.SetCount(0, config.ReqCount)
	profileRepo.SetCount(0, config.ErrorCount)

	return Routes(*profileHandler, *privateKeyHandler, *brokerHandler), nil
}
