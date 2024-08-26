package validator

import (
	"errors"
	"user-management-backend/internal/model"
)

func ValidateUser(user *model.User) error {
	if user.UserName == "" {
		return errors.New("user_name is required")
	}
	if user.FirstName == "" {
		return errors.New("first_name is required")
	}
	if user.LastName == "" {
		return errors.New("last_name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.UserStatus == "" {
		return errors.New("user_status is required")
	}
	return nil
}
