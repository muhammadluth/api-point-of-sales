package repository

import (
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/model"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/muhammadluth/log"
)

type ReferenceRepo struct {
	database *pg.DB
}

func NewReferenceRepo(database *pg.DB) authentication.IReferenceRepo {
	return &ReferenceRepo{database}
}

func (r *ReferenceRepo) ValidationRegisterUserDB(uniqID string,
	request model.RequestCreateUser) (dataUsers []model.TableUsers, err error) {

	err = r.database.Model(&dataUsers).
		Where("username = ?", request.Username).Select()
	if err != nil {
		log.Error(err, uniqID)
		return dataUsers, errors.New("Username is Exist")
	} else if len(dataUsers) != 0 {
		return dataUsers, errors.New("Username is Exist")
	}

	err = r.database.Model(&dataUsers).
		Where("email = ?", request.Email).Select()
	if err != nil {
		log.Error(err, uniqID)
		return dataUsers, errors.New("Email is Exist")
	} else if len(dataUsers) != 0 {
		return dataUsers, errors.New("Email is Exist")
	}

	return dataUsers, nil
}
