package authentication

import (
	"api-point-of-sales/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ILoginUsecase interface {
	Login(ctx *fiber.Ctx) error
}

type IRegisterUsecase interface {
	RegisterUser(ctx *fiber.Ctx) error
}

type IValidationUsecase interface {
	ValidationRegisterUser(uniqID string, request model.RequestCreateUser) error
}

type ICredentialUsecase interface {
	EncryptPassword(uniqID, password string) (string, error)
	VerifyPassword(uniqID, password, passwordHash string) error
}

type ITokenUsecase interface {
	CreateToken(uniqID string, dataUser model.DataUser, ctx *fiber.Ctx) error
	CheckToken(uniqID, accessToken string, ctx *fiber.Ctx) (jwt.StandardClaims, int, error)
}
