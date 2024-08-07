package command

import (
	"errors"
	"regexp"
)

type UserRequest struct {
	Email       string  `json:"email" example:"test@gmail.com"`
	Name        string  `json:"name" example:"test"`
	PhoneNumber *string `json:"phone_number" example:"0123456789"`
}

func (u *UserRequest) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	// email format
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, u.Email); !ok {
		return errors.New("email is invalid")
	}

	if u.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
