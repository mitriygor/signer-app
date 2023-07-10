package broker

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Service interface {
	HandleQueue(entry any)
}

type BrokerService struct {
}

func NewBrokerService() Service {
	return &BrokerService{}
}

func (s *BrokerService) HandleQueue(entry any) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	brokerServiceURL := "http://brocker-service/handle"
	request, err := http.NewRequest("POST", brokerServiceURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return
	}
}
