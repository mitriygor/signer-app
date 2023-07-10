package main

import (
	"fmt"
	"os"
	"strings"
)

type Profile struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Signature    string `json:"signature"`
	Stamp        string `json:"stamp"`
	PrivateKeyID int    `json:"keyId"`
}

func main() {
	file, err := os.Create("75K-100K.psql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	profiles := make([]Profile, 25000)
	for i := range profiles {
		profiles[i] = Profile{
			ID:        i + 75001,
			FirstName: fmt.Sprintf("First%d", i+1+75001),
			LastName:  fmt.Sprintf("Last%d", i+1+75001),
		}
	}

	file.WriteString("INSERT INTO profile (id, first_name, last_name, signature, stamp, private_key_id) VALUES\n")

	values := make([]string, len(profiles))
	for i, profile := range profiles {
		values[i] = fmt.Sprintf("(%d, '%s', '%s', NULL, NULL, NULL)", profile.ID, profile.FirstName, profile.LastName)
	}

	file.WriteString(strings.Join(values, ",\n"))
	file.WriteString(";\n")
}
