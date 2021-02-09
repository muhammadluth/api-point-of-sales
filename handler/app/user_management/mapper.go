package user_management

import "api-point-of-sales/model"

type IUserManagementMapper interface {
	//ROLES
	ToGetRolePayload(dataRoles []model.TableRoles) *[]model.ResponseDropdown
	ToCreateRolePayload(request model.RequestCreateRole) model.TableRoles
	//USERS
	ToGetUsersPayload(dataUsers []model.TableUsers) *[]model.ResponseGetUsers
	ToGetUserByIDPayload(dataUser model.TableUsers) *model.ResponseGetUsers
}

type IReferenceMapper interface {
}
