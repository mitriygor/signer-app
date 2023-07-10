package private_key

import "sync"

type PrivateKey struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Secret string `json:"secret"`
	Mutex  sync.Mutex
}
