package repository

import (
	"api-point-of-sales/handler/app/user_management"

	"github.com/go-pg/pg/v10"
)

type ReferenceRepo struct {
	database *pg.DB
}

func NewReferenceRepo(database *pg.DB) user_management.IReferenceRepo {
	return &ReferenceRepo{database}
}
