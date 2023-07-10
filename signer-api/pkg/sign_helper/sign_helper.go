package sign_helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"log"
	"signer-api/config"
)

func GetStamp() string {
	hash := make([]byte, config.HashSize)
	_, err := rand.Read(hash)
	if err != nil {
		log.Fatal(err)
	}

	stamp := hex.EncodeToString(hash)

	return stamp
}

func Encode(data, secret string) string {

	secretBytes := []byte(secret)

	hasher := hmac.New(md5.New, secretBytes)

	hasher.Write([]byte(data))

	hash := hasher.Sum(nil)

	encodedData := hex.EncodeToString(hash)

	return encodedData
}

func Decode(encodedData, secret string) string {

	secretBytes := []byte(secret)

	hash, err := hex.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}

	hasher := hmac.New(md5.New, secretBytes)

	hasher.Write(hash)

	calculatedHash := hasher.Sum(nil)

	if hmac.Equal(hash, calculatedHash) {

		return string(hash)
	} else {

		panic("Hash verification failed")
	}
}
