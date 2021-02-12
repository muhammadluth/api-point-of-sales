package user_management

import (
	"github.com/gofiber/fiber/v2"
)

type IRoleUsecase interface {
	GetRoles(ctx *fiber.Ctx) error
	CreateRole(ctx *fiber.Ctx) error
}

type IUserUsecase interface {
	GetUsers(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
}

type IValidationUsecase interface {
}

type ICredentialUsecase interface {
}
