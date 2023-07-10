package private_key

import (
	"errors"
	"net/http"
	ah "signer-api/pkg/args_helper"
	jh "signer-api/pkg/json_helper"
)

type Handler struct {
	PrivateKeyService Service
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	queryArgs := r.URL.Query()
	args, err := ah.GetArgs(queryArgs, Args{})

	if err != nil {
		jh.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	privateKeyArgs, ok := args.(Args)
	if !ok {
		jh.ErrorJSON(w, errors.New("Invalid type"), http.StatusBadRequest)
		return
	}

	privateKeys, err := h.PrivateKeyService.GetAllPrivateKeys(privateKeyArgs)
	if err != nil {
		jh.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = jh.WriteJSON(w, http.StatusOK, privateKeys)
}
