package private_key

import "sync"

type PrivateKey struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Secret string `json:"secret"`
	Mutex  sync.Mutex
}

type KeyPayload struct {
	KeyLimit      int `json:"keyLimit,omitempty"`
	BatchSize     int `json:"batchSize,omitempty"`
	WorkersAmount int `json:"workersAmount,omitempty"`
	RecordsAmount int `json:"recordsAmount,omitempty"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Sign   SignPayload `json:"sign,omitempty"`
}

type SignPayload struct {
	Keys          []*PrivateKey `json:"keys,omitempty"`
	KeyLimit      int           `json:"keyLimit,omitempty"`
	BatchSize     int           `json:"batchSize,omitempty"`
	WorkersAmount int           `json:"workersAmount,omitempty"`
	RecordsAmount int           `json:"recordsAmount,omitempty"`
}
