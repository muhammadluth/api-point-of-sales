package authentication

import "api-point-of-sales/model"

type IAuthenticationRepo interface {
	GetUserByUsername(uniqId, account string) (*model.TableUsers, error)
	GetUserByEmail(uniqId, account string) (*model.TableUsers, error)
	InsertUserDB(uniqId string, dataUser model.TableUsers) error
}

type IReferenceRepo interface {
	ValidationRegisterUserDB(uniqId string,
		request model.RequestCreateUser) (dataUsers []model.TableUsers, err error)
}
