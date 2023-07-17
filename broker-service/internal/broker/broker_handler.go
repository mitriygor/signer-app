package broker

import (
	"errors"
	"fmt"
	"net/http"
)
import "broker/pkg/json_helper"

type Handler struct {
	BrokerService Service
}

func (h *Handler) HandleSubmission(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nBroker :: HandleSubmission\n")

	var requestPayload RequestPayload

	err := json_helper.ReadJSON(w, r, &requestPayload)
	if err != nil {
		json_helper.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "log", "key", "sign":
		fmt.Printf("\nBroker :: HandleSubmission :: requestPayload: %v\n", requestPayload)
		h.BrokerService.LogEventViaRabbitService(requestPayload)
	default:
		json_helper.ErrorJSON(w, errors.New("unknown action"))
	}
}
