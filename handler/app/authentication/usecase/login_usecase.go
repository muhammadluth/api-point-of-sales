package usecase

import (
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"

	"github.com/gofiber/fiber/v2"
)

type LoginUsecase struct {
	iAuthenticationMapper authentication.IAuthenticationMapper
	iAuthenticationRepo   authentication.IAuthenticationRepo
	iCredentialUsecase    authentication.ICredentialUsecase
	iTokenUsecase         authentication.ITokenUsecase
}

func NewLoginUsecase(iAuthenticationMapper authentication.IAuthenticationMapper,
	iAuthenticationRepo authentication.IAuthenticationRepo,
	iCredentialUsecase authentication.ICredentialUsecase,
	iTokenUsecase authentication.ITokenUsecase) authentication.ILoginUsecase {
	return &LoginUsecase{iAuthenticationMapper, iAuthenticationRepo, iCredentialUsecase,
		iTokenUsecase}
}

func (u *LoginUsecase) Login(ctx *fiber.Ctx) error {
	var (
		uniqID   = util.CreateUniqID()
		dataUser model.DataUser
		request  model.RequestLogin
	)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Invalid Request Login",
		})
	}
	email, err := u.iAuthenticationRepo.GetUserByEmail(uniqID, request.Account)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Email not found",
		})
	}

	username, err := u.iAuthenticationRepo.GetUserByUsername(uniqID, request.Account)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Username not found",
		})
	}

	if email == nil && username == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "User not found",
		})
	}

	if username != nil && username.Roles != nil {
		dataUser = model.DataUser{
			UserID:      username.UserID,
			FirstName:   username.FirstName,
			LastName:    username.LastName,
			Username:    username.Username,
			Email:       username.Email,
			PhoneNumber: username.PhoneNumber,
			Password:    username.Password,
			Role: model.DataRole{
				RoleName:    username.Roles.RoleName,
				Description: username.Roles.Description,
			},
		}
	} else if email != nil && email.Roles != nil {
		dataUser = model.DataUser{
			UserID:      email.UserID,
			FirstName:   email.FirstName,
			LastName:    email.LastName,
			Username:    email.Username,
			Email:       email.Email,
			PhoneNumber: email.PhoneNumber,
			Password:    email.Password,
			Role: model.DataRole{
				RoleName:    email.Roles.RoleName,
				Description: email.Roles.Description,
			},
		}
	}

	err = u.iCredentialUsecase.VerifyPassword(uniqID, request.Password,
		dataUser.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Invalid Password",
		})
	}

	err = u.iTokenUsecase.CreateToken(uniqID, dataUser, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
			Status:  constant.ERROR,
			Message: "Token not generated",
		})
	}
	return ctx.JSON(model.ResponseLogin{
		FirstName: dataUser.FirstName,
		LastName:  dataUser.LastName,
		Role:      dataUser.Role.RoleName,
	})
}
