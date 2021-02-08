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

func (r *UserManagementRepo) GetRolesDB(uniqID string) (*[]model.TableRoles, int, error) {
	var (
		dataRoles []model.TableRoles
		totalData int
	)

	totalData, err := r.database.Model(&dataRoles).SelectAndCount()
	if err != nil {
		log.Error(err, uniqID)
		return nil, totalData, err
	}
	return &dataRoles, totalData, nil
}

func (r *UserManagementRepo) InsertRoleDB(uniqID string, dataRole model.TableRoles) error {
	_, err := r.database.Model(&dataRole).Insert()
	if err != nil {
		log.Error(err, uniqID)
	}
	return err
}

func (r *UserManagementRepo) GetUsersDB(uniqID string, params model.ParamsUsers) (*[]model.TableUsers,
	int, error) {
	var (
		dataUsers               []model.TableUsers
		totalData, handlingPage int
	)

	if params.Page == 0 {
		handlingPage = 0
	} else {
		handlingPage = (params.Page - 1) * params.Limit
	}

	totalData, err := r.database.Model(&dataUsers).
		Limit(params.Limit).
		Offset(handlingPage).
		Relation("Roles").
		SelectAndCount()
	if err != nil {
		log.Error(err, uniqID)
		return nil, totalData, err
	}
	return &dataUsers, totalData, nil
}
