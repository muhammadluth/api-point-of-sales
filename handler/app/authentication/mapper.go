package authentication

import (
	"api-point-of-sales/model"

	"github.com/dgrijalva/jwt-go"
)

type IAuthenticationMapper interface {
	ToCreateUserPayload(passwordHash string, request model.RequestCreateUser) model.TableUsers
	ToPayloadToken(jwtID, subject, issuer, audience string, issuedAt, expireAt int64) jwt.StandardClaims
}
