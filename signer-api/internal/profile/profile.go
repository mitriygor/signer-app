package profile

import (
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
