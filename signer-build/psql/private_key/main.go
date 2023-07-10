package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type PrivateKey struct {
	ID     int
	Title  string
	Secret string
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func main() {
	file, err := os.Create("insert.psql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	privateKeys := make([]PrivateKey, 100)
	for i := range privateKeys {
		privateKeys[i] = PrivateKey{
			ID:     i + 1,
			Title:  fmt.Sprintf("Key %d", i+1),
			Secret: randomHex(16),
		}
	}

	file.WriteString("INSERT INTO private_key (id, title, secret) VALUES\n")

	values := make([]string, len(privateKeys))
	for i, key := range privateKeys {
		values[i] = fmt.Sprintf("(%d, '%s', '%s')", key.ID, key.Title, key.Secret)
	}

	file.WriteString(strings.Join(values, ",\n"))
	file.WriteString(";\n")
}
