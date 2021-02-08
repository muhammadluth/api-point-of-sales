package usecase

import (
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/model"
	"crypto/rsa"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/muhammadluth/log"
)

type TokenUsecase struct {
	expireAccessToken     time.Duration
	expireRefreshToken    time.Duration
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	iAuthenticationMapper authentication.IAuthenticationMapper
}

func NewTokenUsecase(expireAccessToken time.Duration, expireRefreshToken time.Duration,
	privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey,
	iAuthenticationMapper authentication.IAuthenticationMapper) authentication.ITokenUsecase {
	return &TokenUsecase{expireAccessToken, expireRefreshToken, privateKey, publicKey,
		iAuthenticationMapper}
}

func (u *TokenUsecase) CreateToken(uniqId string, dataUser model.DataUser, ctx *fiber.Ctx) error {
	jwtID, issuer, subject, audiens, issuedAt, expireAccessToken, expireRefreshToken := u.doGeneratePayloadJwt(dataUser, ctx)
	payloadAccessToken := u.iAuthenticationMapper.ToPayloadToken(jwtID, subject, issuer, audiens, issuedAt, expireAccessToken)
	payloadRefreshToken := u.iAuthenticationMapper.ToPayloadToken(jwtID, subject, issuer, audiens, issuedAt, expireRefreshToken)

	strAccessToken, strRefreshToken, err := u.doCreateToken(uniqId, payloadAccessToken, payloadRefreshToken)
	if err != nil {
		log.Error(err, uniqId)
		return err
	}

	now := time.Now().Local()
	ctx.Cookie(&fiber.Cookie{
		Name:     "X-Refresh-Token",
		Value:    strRefreshToken,
		Expires:  now.Add(u.expireRefreshToken),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "lax",
	})
	ctx.Set("X-Access-Token", strAccessToken)
	return err
}

func (u *TokenUsecase) CheckToken(uniqId, accessToken string, ctx *fiber.Ctx) (jwt.StandardClaims, int, error) {
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return u.publicKey, nil
	})
	if err != nil {
		log.Error(err, uniqId)
		return *claims, fiber.StatusUnauthorized, err
	}
	if token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return *claims, fiber.StatusUnauthorized, errors.New("Token Expired")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return *claims, fiber.StatusForbidden, errors.New("Token Has No Access")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return *claims, fiber.StatusUnauthorized, errors.New("Token Invalid")
		} else {
			return *claims, fiber.StatusUnauthorized, errors.New("Token Invalid")
		}
	} else {
		return *claims, fiber.StatusUnauthorized, errors.New("Token Invalid")
	}
	return *claims, fiber.StatusOK, nil
}

func (u *TokenUsecase) doGeneratePayloadJwt(dataUser model.DataUser, ctx *fiber.Ctx) (jwtID,
	issuer, subject, audiens string, issuedAt, expireAccessToken, expireRefreshToken int64) {
	now := time.Now().Local()
	jwtID = strings.Replace(uuid.New().String(), "-", "", -1)
	issuer = ctx.IP()
	subject = dataUser.UserID + ":" + dataUser.Username + ":" + dataUser.Email + ":" + dataUser.Role.RoleName
	audiens = dataUser.Role.RoleName
	issuedAt = now.Unix()
	expireAccessToken = now.Add(u.expireAccessToken).Unix()
	expireRefreshToken = now.Add(u.expireRefreshToken).Unix()
	return jwtID, issuer, subject, audiens, issuedAt, expireAccessToken, expireRefreshToken
}

func (u *TokenUsecase) doCreateToken(uniqId string, payloadAccessToken,
	payloadRefreshToken jwt.StandardClaims) (strAccessToken string, strRefreshToken string, err error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payloadAccessToken)
	strAccessToken, err = accessToken.SignedString(u.privateKey)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payloadRefreshToken)
	strRefreshToken, err = refreshToken.SignedString(u.privateKey)
	return strAccessToken, strRefreshToken, err
}
