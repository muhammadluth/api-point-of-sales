package mapper

import (
	"api-point-of-sales/handler/user_management"
)

type ReferenceMapper struct {
}

func NewReferenceMapper() user_management.IReferenceMapper {
	return &ReferenceMapper{}
}
