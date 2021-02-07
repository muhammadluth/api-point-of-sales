package user_management

import (
	"api-point-of-sales/model"

	"github.com/gofiber/fiber/v2"
)

type IRoleUsecase interface {
	CreateRole(ctx *fiber.Ctx) error
}

type IUserUsecase interface {
	CreateUser(ctx *fiber.Ctx) error
}

type IValidationUsecase interface {
	ValidationCreateUser(traceId string, request model.RequestCreateUser) error
}

type ICredentialUsecase interface {
	EncryptPassword(password string) (string, error)
}
