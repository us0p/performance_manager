package brules

import (
	"context"
	"errors"
	"fmt"
	"performance_manager/src/bRules/repositories"
	"performance_manager/src/entities"
)

type AuthUC struct {
	UsrRp repositories.UserRepo
}

var (
	ErrCreateUser = errors.New("failed to create user")
)

func (auc AuthUC) AuthCreateUser(ctx context.Context, email, name, pass string) (entities.User, error) {
	cu := entities.NewCreateUserEnt(email, name, pass)
	user, err := auc.UsrRp.CreateUser(ctx, cu)
	if err != nil {
		return entities.User{}, fmt.Errorf("%w: %w", ErrCreateUser, err)
	}

	return user, nil
}
