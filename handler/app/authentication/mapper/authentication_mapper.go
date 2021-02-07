package mapper

import (
	"api-point-of-sales/handler/app/authentication"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationMapper struct {
}

func NewAuthenticationMapper() authentication.IAuthenticationMapper {
	return &AuthenticationMapper{}
}

func (m *AuthenticationMapper) ToPayloadToken(jwtID, subject, issuer, audience string, issuedAt,
	expireAt int64) jwt.StandardClaims {

	return jwt.StandardClaims{
		Id:        jwtID,
		Issuer:    issuer,
		Subject:   subject,
		Audience:  audience,
		IssuedAt:  issuedAt,
		ExpiresAt: expireAt,
	}
}
