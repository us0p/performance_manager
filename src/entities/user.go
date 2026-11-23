package entities

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("email is empty")
	ErrInvalidName = errors.New("name is invalid")
)

type CreateUser struct {
	Email    string
	Name     string
	Password string
}

func NewCreateUserEnt(email, name, pass string) CreateUser {
	return CreateUser{
		email,
		name,
		pass,
	}
}

func (cu CreateUser) Validate() []error {
	errs := make([]error, 0)
	if cu.Email == "" {
		errs = append(errs, ErrInvalidEmail)	
	}

	if cu.Name == "" {
		errs = append(errs, ErrInvalidName)
	}

	passREXP := regexp.Compile("a-z+A-Z+0-9+")
	if cu.Password == "" || len(cu.Password) < 8 || 
}

type User struct {
	Email string
	Name  string
	ID    string
}
