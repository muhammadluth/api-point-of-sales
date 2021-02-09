package user_management

import "api-point-of-sales/model"

type IUserManagementRepo interface {
	// ROLES
	GetRolesDB(uniqID string) (*[]model.TableRoles, int, error)
	InsertRoleDB(uniqID string, dataRole model.TableRoles) error
	// USERS
	GetUsersDB(uniqID string, params model.ParamsUsers) (*[]model.TableUsers, int, error)
	GetUserByIDDB(uniqID, paramsID string) (*model.TableUsers, error)
}

type IReferenceRepo interface {
}
