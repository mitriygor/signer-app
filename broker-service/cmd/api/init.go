package main

import (
	"broker/config"
	"broker/internal/broker"
	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

func initHandlers(redisClient *redis.Client, rabbitConn *amqp.Connection) (http.Handler, error) {

	brokerRepo := broker.NewBrokerRepository(redisClient, rabbitConn)
	brokerService := broker.NewBrokerService(brokerRepo)
	brokerHandler := &broker.Handler{BrokerService: brokerService}

	brokerRepo.SetCount(0, config.ReqCount)
	brokerRepo.SetCount(0, config.ErrorCount)

	return Routes(*brokerHandler), nil
}
