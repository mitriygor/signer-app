package broker

import (
	"net/http"
	jh "signer-api/pkg/json_helper"
)

type Handler struct {
	BrokerService Service
}

func (h *Handler) HandleQueueHandler(w http.ResponseWriter, r *http.Request) {

	//payload := RequestPayload{
	//	Action: "log",
	//	Log: LogPayload{
	//		Name:      "log",
	//		Type:      "TEST",
	//		Stamp:     "test_stamp",
	//		Signature: "test_signature",
	//	},
	//}
	//
	//h.BrokerService.HandleQueue(payload)

	_ = jh.WriteJSON(w, http.StatusOK, "Broker handles it!!!")
}
