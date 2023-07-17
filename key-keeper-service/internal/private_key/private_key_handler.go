package private_key

import (
	jh "key-keeper-service/pkg/json_helper"
	"net/http"
)

type Handler struct {
	PrivateKeyService Service
}

func (h *Handler) GetKeysHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload KeyPayload
	var privateKeys []*PrivateKey

	err := jh.ReadJSON(w, r, &requestPayload)
	if err != nil {
		jh.ErrorJSON(w, err)
		return
	}

	privateKeys, err = h.PrivateKeyService.GetPrivateKeys(requestPayload)
	if err != nil {
		jh.ErrorJSON(w, err)
		return
	}

	h.PrivateKeyService.HandleQueue(privateKeys, requestPayload)

	_ = jh.WriteJSON(w, http.StatusOK, "")
}
