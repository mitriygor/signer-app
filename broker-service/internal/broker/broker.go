package broker

type Broker struct {
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

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
