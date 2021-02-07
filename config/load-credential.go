package config

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func LoadCredential(filePrivatePEM, filePublicPEM string) (*rsa.PrivateKey, *rsa.PublicKey) {
	readPrivateKey, err := ioutil.ReadFile(filePrivatePEM)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	readPublicKey, err := ioutil.ReadFile(filePublicPEM)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readPrivateKey)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(readPublicKey)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	return privateKey, publicKey
}
