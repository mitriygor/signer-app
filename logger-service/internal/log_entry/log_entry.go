package log_entry

import "time"

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	Type      string    `bson:"type" json:"type,omitempty"`
	Stamp     string    `bson:"stamp" json:"stamp,omitempty"`
	Signature string    `bson:"signature" son:"signature,omitempty"`
	ProfileID int       `bson:"profile_id" json:"profileID,omitempty"`
	KeyID     int       `bson:"key_id" json:"keyID,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type JSONPayload struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Stamp     string `json:"stamp"`
	Signature string `json:"signature"`
	ProfileID int    `json:"profileID"`
	KeyID     int    `json:"keyID"`
	Data      string `json:"data"`
}
