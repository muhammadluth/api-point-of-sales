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

type IForgetPasswordUsecase interface {
	ForgetPassword(ctx *fiber.Ctx) error
}

type IValidationUsecase interface {
	ValidationRegisterUser(uniqId string, request model.RequestCreateUser) error
}

type ICredentialUsecase interface {
	EncryptPassword(uniqId, password string) (string, error)
	VerifyPassword(uniqId, password, passwordHash string) error
}

type ITokenUsecase interface {
	CreateToken(uniqId string, dataUser model.DataUser, ctx *fiber.Ctx) error
	CheckToken(uniqId, accessToken string, ctx *fiber.Ctx) (jwt.StandardClaims, int, error)
}
