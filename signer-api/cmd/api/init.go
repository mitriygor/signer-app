package main

import (
	"database/sql"
	"net/http"
	"signer-api/internal/broker"
	"signer-api/internal/private_key"
	"signer-api/internal/profile"
)

func initHandlers(conn *sql.DB, connWriter *sql.DB) (http.Handler, error) {
	privateKeyRepo := private_key.NewPrivateKeyRepository(conn)
	privateKeyService := private_key.NewPrivateKeyService(privateKeyRepo)
	privateKeyHandler := &private_key.Handler{PrivateKeyService: privateKeyService}

	brokerService := broker.NewBrokerService()
	brokerHandler := &broker.Handler{BrokerService: brokerService}

	profileRepo := profile.NewProfileRepository(conn, privateKeyRepo, connWriter, brokerService)
	profileService := profile.NewProfileService(profileRepo)
	profileHandler := &profile.Handler{ProfileService: profileService}

	return Routes(*profileHandler, *privateKeyHandler, *brokerHandler), nil
}
