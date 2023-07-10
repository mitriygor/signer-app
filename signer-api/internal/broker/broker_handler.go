package broker

import (
	"net/http"
	jh "signer-api/pkg/json_helper"
)

type Handler struct {
	BrokerService Service
}

func (h *Handler) HandleQueueHandler(w http.ResponseWriter, r *http.Request) {
	h.BrokerService.HandleQueue(nil)

	_ = jh.WriteJSON(w, http.StatusOK, "Broker handles it")
}
