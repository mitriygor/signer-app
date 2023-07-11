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
	var requestPayload RequestPayload

	err := json_helper.ReadJSON(w, r, &requestPayload)
	if err != nil {
		json_helper.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "log":
		fmt.Printf("COUNT: %v\n", h.BrokerService.GetCountService())
		h.BrokerService.LogEventViaRabbitService(requestPayload.Log)
	default:
		json_helper.ErrorJSON(w, errors.New("unknown action"))
	}
}
