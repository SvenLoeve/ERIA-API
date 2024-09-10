package aes

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"
	"viseh-api/database/query"
)

func GenerateKey() string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	return hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
}

func DailyKeyGen() {
	query.SetKey(context.Background(), GenerateKey())
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				query.SetKey(context.Background(), GenerateKey())
			}
		}
	}()
}
