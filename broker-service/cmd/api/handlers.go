package main

import (
	"broker/event"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestPayload struct {
	Action string     `json:"action"`
	Log    LogPayload `json:"log,omitempty"`
}

type LogPayload struct {
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
	Stamp     string `json:"stamp,omitempty"`
	Signature string `json:"signature,omitempty"`
	ProfileID int    `json:"profileID,omitempty"`
	KeyID     int    `json:"keyID,omitempty"`
	Data      string `json:"data,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "log":
		app.logEventViaRabbit(w, requestPayload.Log)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) logEventViaRabbit(w http.ResponseWriter, l LogPayload) {

	err := app.pushToQueue(l)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged via RabbitMQ"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) pushToQueue(l LogPayload) error {
	emitter, err := event.NewEmitter(app.Rabbit)
	if err != nil {
		return err
	}
	defer emitter.Close()

	j, _ := json.MarshalIndent(&l, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		fmt.Printf("\nBrokerService :: pushToQueue: ERROR: %v\n", err.Error())
		return err
	}
	return nil
}
