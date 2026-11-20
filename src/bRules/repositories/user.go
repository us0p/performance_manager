// repository interfaces should be defined in the use cases layer since that's where they're going to be used.
package repositories

import (
	"context"
	"performance_manager/src/entities"
)

type UserRepo interface {
	CreateUser(ctx context.Context, u entities.CreateUser) (entities.User, error)
}
