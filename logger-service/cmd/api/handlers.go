package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Stamp     string `json:"stamp"`
	Signature string `json:"signature"`
	ProfileID int    `json:"profileID"`
	KeyID     int    `json:"keyID"`
	Data      string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name:      requestPayload.Name,
		Type:      requestPayload.Type,
		Stamp:     requestPayload.Stamp,
		Signature: requestPayload.Signature,
		ProfileID: requestPayload.ProfileID,
		KeyID:     requestPayload.KeyID,
		Data:      requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
