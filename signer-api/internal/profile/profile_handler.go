package profile

import (
	"github.com/pkg/errors"
	"net/http"
	ah "signer-api/pkg/args_helper"
	jh "signer-api/pkg/json_helper"
)

type Handler struct {
	ProfileService Service
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	queryArgs := r.URL.Query()
	args, err := ah.GetArgs(queryArgs, Args{})

	if err != nil {
		jh.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	profileArgs, ok := args.(Args)
	if !ok {
		jh.ErrorJSON(w, errors.New("Invalid type"), http.StatusBadRequest)
		return
	}

	profiles, err := h.ProfileService.GetAllProfiles(profileArgs)
	if err != nil {
		jh.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = jh.WriteJSON(w, http.StatusOK, profiles)
}

func (h *Handler) SignAllHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload SignPayload

	err := jh.ReadJSON(w, r, &requestPayload)
	if err != nil {
		jh.ErrorJSON(w, err)
		return
	}

	h.ProfileService.SignAllProfilesWithParams(requestPayload)
	_ = jh.WriteJSON(w, http.StatusOK, "It works")
}
