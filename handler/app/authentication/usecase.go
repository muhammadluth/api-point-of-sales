package authentication

import (
	"api-point-of-sales/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ILoginUsecase interface {
	Login(ctx *fiber.Ctx) error
}

type ICredentialUsecase interface {
	VerifyPassword(traceId, password, passwordHash string) error
}

type ITokenUsecase interface {
	CreateToken(traceId string, dataUser model.DataUser, ctx *fiber.Ctx) error
	CheckToken(traceId, accessToken string, ctx *fiber.Ctx) (jwt.StandardClaims, int, error)
}
