package private_key

import (
	"bytes"
	"encoding/json"
	"fmt"
	"key-keeper-service/config"
	"net/http"
)

type Service interface {
	GetPrivateKeys(args KeyPayload) ([]*PrivateKey, error)
	HandleQueue(privateKeys []*PrivateKey, keyPayload KeyPayload)
}

type privateKeyService struct {
	privateKeyRepo Repository
}

func NewPrivateKeyService(repo Repository) Service {
	return &privateKeyService{
		privateKeyRepo: repo,
	}
}

func (s *privateKeyService) GetPrivateKeys(args KeyPayload) ([]*PrivateKey, error) {
	return s.privateKeyRepo.GetKeys(args)
}

func (s *privateKeyService) HandleQueue(privateKeys []*PrivateKey, keyPayload KeyPayload) {
	signPayload := RequestPayload{
		Action: "sign",
		Sign: SignPayload{
			Keys:          privateKeys,
			KeyLimit:      keyPayload.KeyLimit,
			BatchSize:     keyPayload.BatchSize,
			WorkersAmount: keyPayload.WorkersAmount,
			RecordsAmount: keyPayload.RecordsAmount,
		},
	}

	jsonData, _ := json.MarshalIndent(signPayload, "", "\t")
	request, err := http.NewRequest("POST", config.BrokerService, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("\nSignerAPI::Broker::HandleQueue::ERROR 1:%v\n", err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("\nSignerAPI::Broker::HandleQueue::ERROR 2:%v\n", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return
	}
}
