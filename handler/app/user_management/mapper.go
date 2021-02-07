package user_management

import "api-point-of-sales/model"

type IUserManagementMapper interface {
	ToCreateRolePayload(request model.RequestCreateRole) model.TableRoles
	ToCreateUserPayload(passwordHash string, request model.RequestCreateUser) model.TableUsers
}

type IReferenceMapper interface {
}
