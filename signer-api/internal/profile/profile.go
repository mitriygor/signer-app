package profile

import (
	"signer-api/internal/private_key"
	"sync"
	"time"
)

type Profile struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Signature    string    `json:"signature"`
	Stamp        string    `json:"stamp"`
	PrivateKeyID int       `json:"keyId"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Mutex        sync.Mutex
}

type RequestPayload struct {
	Action string      `json:"action"`
	Sign   SignPayload `json:"sign,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type SignPayload struct {
	Keys          []*private_key.PrivateKey `json:"keys,omitempty"`
	KeyLimit      int                       `json:"keyLimit,omitempty"`
	BatchSize     int                       `json:"batchSize,omitempty"`
	WorkersAmount int                       `json:"workersAmount,omitempty"`
	RecordsAmount int                       `json:"recordsAmount,omitempty"`
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
