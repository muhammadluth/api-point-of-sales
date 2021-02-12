package usecase

import (
	"api-point-of-sales/handler/user_management"
)

type ValidationUsecase struct {
	iReferenceMapper user_management.IReferenceMapper
	iReferenceRepo   user_management.IReferenceRepo
}

func NewValidationUsecase(iReferenceMapper user_management.IReferenceMapper,
	iReferenceRepo user_management.IReferenceRepo) user_management.IValidationUsecase {
	return &ValidationUsecase{iReferenceMapper, iReferenceRepo}
}
