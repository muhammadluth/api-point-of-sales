package usecase

import (
	"api-point-of-sales/handler/user_management"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"

	"github.com/gofiber/fiber/v2"
)

type RoleUsecase struct {
	iUserManagementMapper user_management.IUserManagementMapper
	iUserManagementRepo   user_management.IUserManagementRepo
}

func NewRoleUsecase(iUserManagementMapper user_management.IUserManagementMapper,
	iUserManagementRepo user_management.IUserManagementRepo) user_management.IRoleUsecase {
	return &RoleUsecase{iUserManagementMapper, iUserManagementRepo}
}

func (u *RoleUsecase) GetRoles(ctx *fiber.Ctx) error {
	var uniqId = util.CreateUniqID()

	dataRoles, totalData, err := u.iUserManagementRepo.GetRolesDB(uniqId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Retrieve Data Role",
		})
	}
	response := u.iUserManagementMapper.ToGetRolePayload(*dataRoles)
	return ctx.JSON(model.ResponseSuccessWithoutPagination{
		TotalData: totalData,
		Data:      *response,
	})
}

func (u *RoleUsecase) CreateRole(ctx *fiber.Ctx) error {
	var (
		uniqId  = ctx.Locals("uniqId").(string)
		request model.RequestCreateRole
	)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Invalid Request Create Role",
		})
	}
	dataRole := u.iUserManagementMapper.ToCreateRolePayload(request)
	if err := u.iUserManagementRepo.InsertRoleDB(uniqId, dataRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Error Create Role",
		})
	}
	return ctx.JSON(model.ResponseHTTP{
		Status:  constant.SUCCESS,
		Message: "Successfully Create Role",
	})
}
