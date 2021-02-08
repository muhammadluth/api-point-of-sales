package usecase

import (
	"api-point-of-sales/handler/app/authentication"
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

func NewCredentialUsecase(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) authentication.ICredentialUsecase {
	return &CredentialUsecase{privateKey, publicKey}
}

func (u *CredentialUsecase) EncryptPassword(uniqID, password string) (string, error) {
	passwordHash := sha256.Sum256([]byte(password))
	signaturePassword, err := rsa.SignPSS(rand.Reader, u.privateKey, crypto.SHA256, passwordHash[:], nil)
	if err != nil {
		log.Error(err, uniqID)
		return "", err
	}
	return hex.EncodeToString(signaturePassword), nil
}

func (u *CredentialUsecase) VerifyPassword(uniqID, password, passwordHash string) error {
	hash := sha256.Sum256([]byte(password))
	passwordSignature, err := hex.DecodeString(passwordHash)
	if err != nil {
		log.Error(err, uniqID)
		return err
	}
	return rsa.VerifyPSS(u.publicKey, crypto.SHA256, hash[:], passwordSignature, nil)
}
