package user_management

import "api-point-of-sales/model"

type IUserManagementRepo interface {
	GetRolesDB(uniqID string) (*[]model.TableRoles, int, error)
	InsertRoleDB(uniqID string, dataRole model.TableRoles) error
	GetUsersDB(uniqID string, params model.ParamsUsers) (*[]model.TableUsers, int, error)
}

type IReferenceRepo interface {
}
