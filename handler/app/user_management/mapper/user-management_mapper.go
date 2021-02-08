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

func (m *UserManagementMapper) ToGetRolePayload(dataRoles []model.TableRoles) *[]model.ResponseDropdown {
	var response []model.ResponseDropdown
	for i := range dataRoles {
		payload := model.ResponseDropdown{
			Label: strings.Title(dataRoles[i].RoleName),
			Value: dataRoles[i].RoleID,
		}
		response = append(response, payload)
	}
	return &response
}

func (m *UserManagementMapper) ToCreateRolePayload(request model.RequestCreateRole) model.TableRoles {
	return model.TableRoles{
		RoleID:      strings.Replace(uuid.New().String(), "-", "", -1),
		RoleName:    strings.Title(request.RoleName),
		Description: strings.Title(request.Description),
	}
}

func (m *UserManagementMapper) ToGetUsersPayload(dataUsers []model.TableUsers) *[]model.ResponseGetUsers {
	var response []model.ResponseGetUsers
	for i := range dataUsers {
		payload := model.ResponseGetUsers{
			UserID:      dataUsers[i].UserID,
			FirstName:   dataUsers[i].FirstName,
			LastName:    dataUsers[i].LastName,
			Username:    dataUsers[i].Username,
			Email:       dataUsers[i].Email,
			PhoneNumber: dataUsers[i].PhoneNumber,
			Role:        dataUsers[i].Roles.RoleName,
		}
		response = append(response, payload)
	}
	return &response
}
