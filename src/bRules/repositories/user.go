package repositories

import (
	"context"
	"performance_manager/src/entities"
)

type UserRepo interface {
	CreateUser(ctx context.Context, u entities.CreateUser) (entities.User, error)
}
