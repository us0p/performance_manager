package controllers

import (
	"context"
	brules "performance_manager/src/bRules"
)

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
