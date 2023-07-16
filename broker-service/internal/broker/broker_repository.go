package broker

import (
	"broker/config"
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
	IncrCount(countName string)
	GetCount(countName string) int
	SetCount(count int, countName string)
}

type BrokerRepository struct {
	redisClient *redis.Client
	rabbitConn  *amqp.Connection
	emitter     *event.Emitter
}

func NewBrokerRepository(redisClient *redis.Client, rabbitConn *amqp.Connection) Repository {

	emitter, err := event.NewEmitter(rabbitConn)

	if err != nil {
		fmt.Printf("\nNewBrokerRepository :: ERROR: %v\n", err.Error())
		return nil
		//panic(err) // Handle error appropriately
	}

	return &BrokerRepository{
		redisClient: redisClient,
		rabbitConn:  rabbitConn,
		emitter:     emitter,
	}
}

func (b *BrokerRepository) LogEventViaRabbit(l LogPayload) {
	err := b.PushToQueue(l)
	if err != nil {
		fmt.Printf("\nBrokerRepo :: LogEventViaRabbit :: err:%v\n", err.Error())
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "logged via RabbitMQ"
}

func (b *BrokerRepository) PushToQueue(l LogPayload) error {
	j, _ := json.MarshalIndent(&l, "", "\t")
	err := b.emitter.Push(string(j), "log.INFO")
	if err != nil {
		b.IncrCount(config.ErrorCount)
		fmt.Printf("\nBrokerService :: PushToQueue: error: %v\n", err.Error())
		return err
	}

	b.IncrCount(config.ReqCount)

	return nil
}

func (b *BrokerRepository) IncrCount(countName string) {
	b.redisClient.Incr(context.TODO(), countName)
}

func (b *BrokerRepository) SetCount(count int, countName string) {
	b.redisClient.Set(context.TODO(), countName, count, 0)
}

func (b *BrokerRepository) GetCount(countName string) int {
	count, err := b.redisClient.Get(context.TODO(), countName).Int()
	if err != nil {
		fmt.Printf("\nBroker :: Redis :: GetCount :: ERROR:%v\n", err.Error())
		return -1
	}

	return count
}
