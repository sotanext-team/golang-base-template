package utils

import (
	"errors"

	"app-api/modules/auth/custom_types/request_input"

	"github.com/asaskevich/govalidator"
)

func ValidateUserInput(user request_input.UserInput) error {
	if govalidator.IsNull(user.Username) || govalidator.IsNull(user.Email) || govalidator.IsNull(user.Password) {
		return errors.New("Data cannot be empty")
	}

	if !govalidator.IsEmail(user.Email) {
		return errors.New("Email is invalid")
	}
	return nil
}

func ValidateUserLogin(user request_input.UserInput) error {
	if (govalidator.IsNull(user.Username) && govalidator.IsNull(user.Email)) || govalidator.IsNull(user.Password) {
		return errors.New("Data cannot be empty")
	}

	if !govalidator.IsEmail(user.Email) {
		return errors.New("Email is invalid")
	}
	return nil
}
