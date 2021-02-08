package usecase

import (
	"api-point-of-sales/handler/app/user_management"
	"crypto/rsa"
)

type CredentialUsecase struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewCredentialUsecase(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) user_management.ICredentialUsecase {
	return &CredentialUsecase{privateKey, publicKey}
}
