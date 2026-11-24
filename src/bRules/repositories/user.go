// repository interfaces should be defined in the use cases layer since that's where they're going to be used.
package repositories

import (
	"context"
	"errors"
	"performance_manager/src/entities"
)

// We define the error close to the source of it for better readability.
var (
	ErrMailTaken = errors.New("email already registered")
)

type UserRepo interface {
	CreateUser(ctx context.Context, u entities.CreateUser) (entities.User, error)
}
