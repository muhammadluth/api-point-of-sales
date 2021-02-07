package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

type IAuthenticationMapper interface {
	ToPayloadToken(jwtID, subject, issuer, audience string, issuedAt, expireAt int64) jwt.StandardClaims
}
