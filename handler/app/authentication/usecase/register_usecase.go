package usecase

import (
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"

	"github.com/gofiber/fiber/v2"
)

type RegisterUsecase struct {
	iAuthenticationMapper authentication.IAuthenticationMapper
	iAuthenticationRepo   authentication.IAuthenticationRepo
	iCredentialUsecase    authentication.ICredentialUsecase
	iValidationUsecase    authentication.IValidationUsecase
}

func NewRegisterUsecase(iAuthenticationMapper authentication.IAuthenticationMapper,
	iAuthenticationRepo authentication.IAuthenticationRepo,
	iCredentialUsecase authentication.ICredentialUsecase,
	iValidationUsecase authentication.IValidationUsecase) authentication.IRegisterUsecase {
	return &RegisterUsecase{iAuthenticationMapper, iAuthenticationRepo, iCredentialUsecase,
		iValidationUsecase}
}

func (u *RegisterUsecase) RegisterUser(ctx *fiber.Ctx) error {
	var (
		uniqID  = util.CreateUniqID()
		request model.RequestCreateUser
	)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Invalid Request Register User",
		})
	}

	if errValid := u.iValidationUsecase.ValidationRegisterUser(uniqID, request); errValid != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: errValid.Error(),
		})
	}

	passwordHash, err := u.iCredentialUsecase.EncryptPassword(uniqID, request.ConfirmPassword)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Generate Password",
		})
	} else if passwordHash == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Generate Password",
		})
	}

	dataUser := u.iAuthenticationMapper.ToCreateUserPayload(passwordHash, request)
	if err := u.iAuthenticationRepo.InsertUserDB(uniqID, dataUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Register User",
		})
	}

	return ctx.JSON(model.ResponseHTTP{
		Status:  constant.SUCCESS,
		Message: "Successfully Register User",
	})
}
