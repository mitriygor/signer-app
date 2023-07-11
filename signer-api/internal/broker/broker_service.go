package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"signer-api/config"
)

type Service interface {
	HandleQueue(payload RequestPayload)
}

type BrokerService struct {
}

func NewBrokerService() Service {
	return &BrokerService{}
}

func (s *BrokerService) HandleQueue(payload RequestPayload) {
	jsonData, _ := json.MarshalIndent(payload, "", "\t")
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
