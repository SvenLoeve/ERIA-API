package aes

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"viseh-api/database/query"
)

func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(decryptKey(keyString))
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}

func decryptKey(encryptedString string) (decryptedString string) {
	keys, err := query.GetAllKeys(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	for _, crypto := range keys {

		key, _ := hex.DecodeString(crypto.Key)
		enc, _ := hex.DecodeString(encryptedString)

		//Create a new Cipher Block from the key
		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		//Create a new GCM
		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		//Get the nonce size
		nonceSize := aesGCM.NonceSize()

		//Extract the nonce from the encrypted data
		nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

		//Decrypt the data
		plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			println(err.Error())
		} else {
			return string(plaintext)
		}
	}
	return ""
}
