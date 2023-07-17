package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"listener/config"
	"listener/internal/listener"
	"net/http"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn         *amqp.Connection
	queueName    string
	listenerRepo listener.Repository
}

func NewConsumer(conn *amqp.Connection, listenerRepo listener.Repository) (Consumer, error) {
	consumer := Consumer{
		conn:         conn,
		listenerRepo: listenerRepo,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		fmt.Printf("Error getting channel: %v\n", err.Error())
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		fmt.Printf("Error declaring queue: %v\n", err.Error())
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)

		if err != nil {
			fmt.Printf("Error binding queue: %v\n", err.Error())
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("Error consuming queue: %v\n", err.Error())
		return err
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for d := range messages {
				var payload listener.Payload
				_ = json.Unmarshal(d.Body, &payload)
				consumer.HandlePayload(payload)
			}
		}()
	}
	wg.Wait()

	return nil
}

func (consumer *Consumer) HandlePayload(payload listener.Payload) {
	switch payload.Action {
	case "key":

		for _, c := range config.Counts {
			consumer.listenerRepo.SetCount(0, c)
		}

		err := getKeys(payload.Key)
		if err != nil {
			fmt.Printf("handlePayload::key::ERROR:%v\n", err.Error())
			consumer.listenerRepo.IncrCount(config.ErrorCount)
		} else {
			consumer.listenerRepo.IncrCount(config.ReqCount)
		}
	case "sign":
		err := getSigns(payload.Sign)
		if err != nil {
			fmt.Printf("handlePayload::sign::ERROR:%v\n", err.Error())
			consumer.listenerRepo.IncrCount(config.ErrorCount)
		} else {
			consumer.listenerRepo.IncrCount(config.ReqCount)
		}
	case "log":
		err := logEvent(payload.Log)
		if err != nil {
			fmt.Printf("handlePayload::log::ERROR:%v\n", err.Error())
			consumer.listenerRepo.IncrCount(config.ErrorCount)
		} else {
			consumer.listenerRepo.IncrCount(config.ReqCount)
		}
	default:
		fmt.Printf("handlePayload::ERROR:%v\n", "WRONG TYPE")
	}
}

func getKeys(entry listener.KeyPayload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	keyKeeperServiceURL := "http://key-keeper-service/keys"

	request, err := http.NewRequest("POST", keyKeeperServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("\nConsumer :: getKeys :: ERROR 1: %v\n", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("\nConsumer :: getKeys :: ERROR 2: %v\n", err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

func getSigns(entry listener.SignPayload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	keyKeeperServiceURL := "http://signer-api/sign"

	request, err := http.NewRequest("POST", keyKeeperServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("\nConsumer :: getKeys :: ERROR 1: %v\n", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("\nConsumer :: getKeys :: ERROR 2: %v\n", err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

func logEvent(entry listener.LogPayload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("logEvent::ERROR 1:%v\n", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("logEvent::ERROR 2:%v\n", err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
