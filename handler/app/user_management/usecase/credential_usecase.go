package usecase

import (
	"api-point-of-sales/handler/app/user_management"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"

	"github.com/muhammadluth/log"
)

type CredentialUsecase struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewCredentialUsecase(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) user_management.ICredentialUsecase {
	return &CredentialUsecase{privateKey, publicKey}
}

func (u *CredentialUsecase) EncryptPassword(password string) (string, error) {
	passwordHash := sha256.Sum256([]byte(password))
	signaturePassword, err := rsa.SignPSS(rand.Reader, u.privateKey, crypto.SHA256, passwordHash[:], nil)
	if err != nil {
		log.Error(err, err.Error())
		return "", err
	}
	return hex.EncodeToString(signaturePassword), nil
}
