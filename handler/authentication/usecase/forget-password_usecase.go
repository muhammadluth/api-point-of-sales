package usecase

import (
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"

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
	return ctx.JSON(model.ResponseHTTP{
		Status:  constant.SUCCESS,
		Message: "Sorry, This Feature Is Currently Under Development!",
	})
}
