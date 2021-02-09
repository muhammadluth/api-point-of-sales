package usecase

import (
	"api-point-of-sales/handler/app/authentication"

	"github.com/gofiber/fiber/v2"
)

type ForgetPasswordUsecase struct {
	iAuthenticationMapper authentication.IAuthenticationMapper
	iAuthenticationRepo   authentication.IAuthenticationRepo
	iCredentialUsecase    authentication.ICredentialUsecase
}

func NewForgetPasswordUsecase(iAuthenticationMapper authentication.IAuthenticationMapper,
	iAuthenticationRepo authentication.IAuthenticationRepo,
	iCredentialUsecase authentication.ICredentialUsecase) authentication.IForgetPasswordUsecase {
	return &ForgetPasswordUsecase{iAuthenticationMapper, iAuthenticationRepo, iCredentialUsecase}
}

func (u *ForgetPasswordUsecase) ForgetPassword(ctx *fiber.Ctx) error {
	return nil
}
