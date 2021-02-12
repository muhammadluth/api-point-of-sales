package usecase

import (
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/model"
	"errors"
	"regexp"
)

type ValidationUsecase struct {
	iReferenceRepo authentication.IReferenceRepo
}

func NewValidationUsecase(iReferenceRepo authentication.IReferenceRepo) authentication.IValidationUsecase {
	return &ValidationUsecase{iReferenceRepo}
}

func (u *ValidationUsecase) ValidationRegisterUser(uniqId string,
	request model.RequestCreateUser) error {
	if isValid, err := u.iReferenceRepo.ValidationRegisterUserDB(uniqId, request); err != nil {
		return err
	} else if len(isValid) != 0 {
		return err
	}
	if err := u.doValidationUsername(request.Username); err != nil {
		return err
	}
	if err := u.doValidationEmail(request.Email); err != nil {
		return err
	}
	if err := u.doValidationPassword(request.Password, request.ConfirmPassword); err != nil {
		return err
	}
	return nil
}

func (u *ValidationUsecase) doValidationUsername(username string) error {
	if len(username) < 5 {
		return errors.New("Username less than 5 character")
	}
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(A_Z, username); b || err != nil {
		return errors.New("Username must lowercase character")
	}
	return nil
}

func (u *ValidationUsecase) doValidationEmail(email string) error {
	regexEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if isValid := regexEmail.MatchString(email); !isValid {
		return errors.New("Email is invalid")
	}
	return nil
}

func (u *ValidationUsecase) doValidationPassword(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("Password and confirm password are different")
	}
	if len(password) < 8 || len(confirmPassword) < 8 {
		return errors.New("Password less than 8 character")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, confirmPassword); !b || err != nil {
		return errors.New("Password need number")
	}
	if b, err := regexp.MatchString(a_z, confirmPassword); !b || err != nil {
		return errors.New("Password need lowercase 1 character")
	}
	if b, err := regexp.MatchString(A_Z, confirmPassword); !b || err != nil {
		return errors.New("Password need uppercase 1 character")
	}
	if b, err := regexp.MatchString(symbol, confirmPassword); !b || err != nil {
		return errors.New("Password need 1 symbol")
	}
	return nil
}
