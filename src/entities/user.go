package entities

import (
	"regexp"
)

var (
	hasLower = regexp.MustCompile(`[a-z]+`)
	hasUpper = regexp.MustCompile(`[A-Z]+`)
	hasDigit = regexp.MustCompile(`[0-9]+`)
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

func (cu CreateUser) Validate() error {
	var err ValidationError
	if cu.Email == "" {
		err.add("email", "can't be empty")
	}

	if cu.Name == "" {
		err.add("name", "can't be empty")
	}

	if !hasLower.Match([]byte(cu.Password)) {
		err.add("password", "must contain lowercase character")
	}

	if !hasUpper.Match([]byte(cu.Password)) {
		err.add("password", "must contain uppercase character")
	}

	if !hasDigit.Match([]byte(cu.Password)) {
		err.add("password", "must contain a digit")
	}

	return err
}

type User struct {
	Email string
	Name  string
	ID    string
}
