package broker

import "broker/internal/private_key"

type Broker struct {
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Log    LogPayload  `json:"log,omitempty"`
	Key    KeyPayload  `json:"key,omitempty"`
	Sign   SignPayload `json:"sign,omitempty"`
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

type KeyPayload struct {
	KeyLimit      int `json:"keyLimit,omitempty"`
	BatchSize     int `json:"batchSize,omitempty"`
	WorkersAmount int `json:"workersAmount,omitempty"`
	RecordsAmount int `json:"recordsAmount,omitempty"`
}

type SignPayload struct {
	Keys          []private_key.PrivateKey `json:"keys,omitempty"`
	KeyLimit      int                      `json:"keyLimit,omitempty"`
	BatchSize     int                      `json:"batchSize,omitempty"`
	WorkersAmount int                      `json:"workersAmount,omitempty"`
	RecordsAmount int                      `json:"recordsAmount,omitempty"`
}
