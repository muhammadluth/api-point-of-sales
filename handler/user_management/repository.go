package user_management

import "api-point-of-sales/model"

type IUserManagementRepo interface {
	// ROLES
	GetRolesDB(uniqId string) (*[]model.TableRoles, int, error)
	InsertRoleDB(uniqId string, dataRole model.TableRoles) error
	// USERS
	GetUsersDB(uniqId string, params model.ParamsUsers) (*[]model.TableUsers, int, error)
	GetUserByIDDB(uniqId, paramsID string) (*model.TableUsers, error)
}

type IReferenceRepo interface {
}
