package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Decrypt(key, data string) ([]byte, error) {
	keyAsByte := []byte(key)

	dataAsByte, errCovertData := hex.DecodeString(data)
	if errCovertData != nil {
		return nil, errCovertData
	}
	blockKey, errBlockKey := aes.NewCipher(keyAsByte)
	if errBlockKey != nil {
		return nil, errBlockKey
	}
	gcm, errGCM := cipher.NewGCM(blockKey)
	if errGCM != nil {
		return nil, errGCM
	}
	nonce, encryptData := dataAsByte[:gcm.NonceSize()], dataAsByte[gcm.NonceSize():]
	decryptData, err := gcm.Open(nil, nonce, encryptData, nil)
	if err != nil {
		return nil, err
	}
	return decryptData, nil
}
