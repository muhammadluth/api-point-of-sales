package usecase

import (
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadluth/log"
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

func (u *UserUsecase) GetUsers(ctx *fiber.Ctx) error {
	var (
		uniqID = util.CreateUniqID()
		params = new(model.ParamsUsers)
	)
	if err := ctx.QueryParser(params); err != nil {
		log.Error(err, uniqID)
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Query Parameters",
		})
	}
	if strings.Title(ctx.Locals("role").(string)) != strings.Title(constant.ROLE_ADMIN) {
		return ctx.Status(fiber.StatusForbidden).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Your Role Has No Access",
		})
	}
	dataUsers, totalData, err := u.iUserManagementRepo.GetUsersDB(uniqID, *params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Retrieve Data Users",
		})
	}
	response := u.iUserManagementMapper.ToGetUsersPayload(*dataUsers)
	return ctx.JSON(util.ResponseSuccessWithPagination(float64(totalData), float64(params.Limit),
		float64(params.Page), response))
}

func (u *UserUsecase) GetUserByID(ctx *fiber.Ctx) error {
	var (
		uniqID   = util.CreateUniqID()
		paramsID = ctx.Params("id")
	)
	if strings.Title(ctx.Locals("role").(string)) != strings.Title(constant.ROLE_ADMIN) &&
		paramsID != ctx.Locals("user_id").(string) {
		return ctx.Status(fiber.StatusForbidden).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Your Role Has No Access",
		})
	}
	dataUser, err := u.iUserManagementRepo.GetUserByIDDB(uniqID, paramsID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Retrieve Data User",
		})
	}
	response := u.iUserManagementMapper.ToGetUserByIDPayload(*dataUser)
	return ctx.JSON(response)
}
