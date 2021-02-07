package user_management

import "api-point-of-sales/model"

type IUserManagementRepo interface {
	InsertRoleDB(traceId string, dataRole model.TableRoles) error
	InsertUserDB(traceId string, dataUser model.TableUsers) error
}

type IReferenceRepo interface {
	ValidationCreateUserDB(traceID string,
		request model.RequestCreateUser) (dataUsers []model.TableUsers, err error)
}
