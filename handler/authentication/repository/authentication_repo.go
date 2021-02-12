package repository

import (
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/model"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/muhammadluth/log"
)

type AuthenticationRepo struct {
	database *pg.DB
}

func NewAuthenticationRepo(database *pg.DB) authentication.IAuthenticationRepo {
	return &AuthenticationRepo{database}
}

func (r *AuthenticationRepo) GetUserByUsername(uniqId, account string) (*model.TableUsers, error) {
	var dataUsers model.TableUsers

	err := r.database.Model(&dataUsers).Where("username = ?", account).
		Relation("Roles").Select()
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		log.Error(err, uniqId)
		return nil, err
	}
	return &dataUsers, nil
}

func (r *AuthenticationRepo) GetUserByEmail(uniqId, account string) (*model.TableUsers, error) {
	var dataUsers model.TableUsers

	err := r.database.Model(&dataUsers).Where("email = ?", account).
		Relation("Roles").Select()
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		log.Error(err, uniqId)
		return nil, err
	}
	return &dataUsers, nil
}

func (r *AuthenticationRepo) InsertUserDB(uniqId string, dataUser model.TableUsers) error {
	_, err := r.database.Model(&dataUser).Insert()
	if err != nil {
		log.Error(err, uniqId)
	}
	return err
}
