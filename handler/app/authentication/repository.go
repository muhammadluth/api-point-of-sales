package authentication

import "api-point-of-sales/model"

type IAuthenticationRepo interface {
	GetUserByUsername(uniqID, account string) (*model.TableUsers, error)
	GetUserByEmail(uniqID, account string) (*model.TableUsers, error)
	InsertUserDB(uniqID string, dataUser model.TableUsers) error
}

type IReferenceRepo interface {
	ValidationRegisterUserDB(uniqID string,
		request model.RequestCreateUser) (dataUsers []model.TableUsers, err error)
}
