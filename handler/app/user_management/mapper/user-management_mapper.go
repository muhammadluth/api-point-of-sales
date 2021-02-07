package mapper

import (
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/model"
	"strings"

	"github.com/google/uuid"
)

type UserManagementMapper struct {
}

func NewUserManagementMapper() user_management.IUserManagementMapper {
	return &UserManagementMapper{}
}

func (m *UserManagementMapper) ToCreateRolePayload(request model.RequestCreateRole) model.TableRoles {
	return model.TableRoles{
		RoleID:      strings.Replace(uuid.New().String(), "-", "", -1),
		RoleName:    strings.Title(request.RoleName),
		Description: strings.Title(request.Description),
	}
}

func (m *UserManagementMapper) ToCreateUserPayload(passwordHash string,
	request model.RequestCreateUser) model.TableUsers {
	return model.TableUsers{
		UserID:      strings.Replace(uuid.New().String(), "-", "", -1),
		FirstName:   strings.Title(request.FirstName),
		LastName:    strings.Title(request.LastName),
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    passwordHash,
		RoleID:      request.RoleID,
	}
}
