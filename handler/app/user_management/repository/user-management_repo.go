package repository

import (
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/model"

	"github.com/go-pg/pg/v10"
	"github.com/muhammadluth/log"
)

type UserManagementRepo struct {
	database *pg.DB
}

func NewUserManagementRepo(database *pg.DB) user_management.IUserManagementRepo {
	return &UserManagementRepo{database}
}

func (r *UserManagementRepo) InsertRoleDB(traceId string, dataRole model.TableRoles) error {
	_, err := r.database.Model(&dataRole).Insert()
	if err != nil {
		log.Error(err, traceId)
	}
	return err
}

func (r *UserManagementRepo) InsertUserDB(traceId string, dataUser model.TableUsers) error {
	_, err := r.database.Model(&dataUser).Insert()
	if err != nil {
		log.Error(err, traceId)
	}
	return err
}
