package broker

import (
	"broker/event"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	LogEventViaRabbit(l LogPayload)
	PushToQueue(l LogPayload) error
	IncrCount()
	GetCount() int
}

type BrokerRepository struct {
	redisClient *redis.Client
	rabbitConn  *amqp.Connection
}

func NewBrokerRepository(redisClient *redis.Client, rabbitConn *amqp.Connection) Repository {
	return &BrokerRepository{
		redisClient: redisClient,
		rabbitConn:  rabbitConn,
	}
}

func (b *BrokerRepository) LogEventViaRabbit(l LogPayload) {
	err := b.PushToQueue(l)
	if err != nil {
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "logged via RabbitMQ"
}

func (b *BrokerRepository) PushToQueue(l LogPayload) error {
	emitter, err := event.NewEmitter(b.rabbitConn)
	if err != nil {
		return err
	}
	defer emitter.Close()

	j, _ := json.MarshalIndent(&l, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		fmt.Printf("\nBrokerService :: pushToQueue: ERROR: %v\n", err.Error())
		return err
	}
	return nil
}

func (b *BrokerRepository) IncrCount() {
	b.redisClient.Incr(context.TODO(), "request_count")
}

func (b *BrokerRepository) GetCount() int {
	count, err := b.redisClient.Get(context.TODO(), "request_count").Int()
	if err != nil {
		fmt.Printf("\nERedis :: GetCount :: ERROR:%v\n", err.Error())
		return -1
	}

	return count
}
