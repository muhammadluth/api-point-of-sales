package usecase

import (
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"

	"github.com/gofiber/fiber/v2"
)

type UserUsecase struct {
	iUserManagementMapper user_management.IUserManagementMapper
	iUserManagementRepo   user_management.IUserManagementRepo
	iCredentialUsecase    user_management.ICredentialUsecase
	iValidationUsecase    user_management.IValidationUsecase
}

func NewUserUsecase(iUserManagementMapper user_management.IUserManagementMapper,
	iUserManagementRepo user_management.IUserManagementRepo,
	iCredentialUsecase user_management.ICredentialUsecase,
	iValidationUsecase user_management.IValidationUsecase) user_management.IUserUsecase {
	return &UserUsecase{iUserManagementMapper, iUserManagementRepo, iCredentialUsecase,
		iValidationUsecase}
}

func (u *UserUsecase) CreateUser(ctx *fiber.Ctx) error {
	var (
		traceId = util.CreateTraceID()
		request model.RequestCreateUser
	)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Invalid Request Create User",
		})
	}

	if errValid := u.iValidationUsecase.ValidationCreateUser(traceId, request); errValid != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: errValid.Error(),
		})
	}

	passwordHash, err := u.iCredentialUsecase.EncryptPassword(request.ConfirmPassword)
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

	dataUser := u.iUserManagementMapper.ToCreateUserPayload(passwordHash, request)
	if err := u.iUserManagementRepo.InsertUserDB(traceId, dataUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Create User",
		})
	}

	return ctx.JSON(model.ResponseHTTP{
		Status:  constant.SUCCESS,
		Message: "Successfully Create User",
	})
}
