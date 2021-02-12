package mapper

import (
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthenticationMapper struct {
}

func NewAuthenticationMapper() authentication.IAuthenticationMapper {
	return &AuthenticationMapper{}
}

func (m *AuthenticationMapper) ToCreateUserPayload(passwordHash string,
	request model.RequestCreateUser) model.TableUsers {
	return model.TableUsers{
		UserID:      strings.Replace(uuid.New().String(), "-", "", -1),
		FirstName:   strings.Title(request.FirstName),
		LastName:    strings.Title(request.LastName),
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    passwordHash,
		RoleID:      request.RoleID,
	}
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
