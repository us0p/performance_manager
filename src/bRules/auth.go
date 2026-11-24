package brules

import (
	"context"
	"performance_manager/src/bRules/interfaces"
	"performance_manager/src/bRules/repositories"
	"performance_manager/src/entities"
)

type AuthUC struct {
	UsrRp  repositories.UserRepo
	Hasher interfaces.Hasher
}

func (auc AuthUC) AuthCreateUser(ctx context.Context, email, name, pass string) (entities.User, error) {
	var u entities.User
	cu := entities.NewCreateUserEnt(email, name, pass)
	err := cu.Validate()
	if err != nil {
		return u, err
	}

	hashedPass, err := auc.Hasher.Hash(cu.Password)
	if err != nil {
		return u, err
	}
	cu.Password = hashedPass

	u, err = auc.UsrRp.CreateUser(ctx, cu)
	if err != nil {
		return u, err
	}

	return u, nil
}
