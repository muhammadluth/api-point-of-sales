package repository

import (
	"api-point-of-sales/handler/app/authentication"
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

func (r *AuthenticationRepo) GetUserByUsername(traceId,
	account string) (*model.TableUsers, error) {
	var dataUsers model.TableUsers

	err := r.database.Model(&dataUsers).Where("username = ?", account).
		Relation("Roles").Select()
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		log.Error(err, traceId)
		return nil, err
	}
	return &dataUsers, nil
}

func (r *AuthenticationRepo) GetUserByEmail(traceId,
	account string) (*model.TableUsers, error) {
	var dataUsers model.TableUsers

	err := r.database.Model(&dataUsers).Where("email = ?", account).
		Relation("Roles").Select()
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		log.Error(err, traceId)
		return nil, err
	}
	return &dataUsers, nil
}
