package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(key, data string) ([]byte, error) {
	keyAsByte := []byte(key)
	dataAsByte := []byte(data)

	blockKey, errBlockKey := aes.NewCipher(keyAsByte)
	if errBlockKey != nil {
		return nil, errBlockKey
	}
	gcm, errGCM := cipher.NewGCM(blockKey)
	if errGCM != nil {
		return nil, errGCM
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, errGCM = rand.Read(nonce); errGCM != nil {
		return nil, errGCM
	}
	encryptData := gcm.Seal(nonce, nonce, dataAsByte, nil)
	return encryptData, nil
}
