package controllers

import (
	"context"
	brules "performance_manager/src/bRules"
)

// The controler should take clean inputs and have no reference to frameworks.
// It must return a presentation of the data ready to be displayed.
type AuthController struct {
	AuthUC brules.AuthUC
}

func (ac AuthController) CreateUserController(ctx context.Context, email, name, pass string) error {
	_, err := ac.AuthUC.AuthCreateUser(ctx, email, name, pass)
	if err != nil {
		return err
	}

	return nil
}
