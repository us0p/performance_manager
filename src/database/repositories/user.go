package repositories

import (
	"context"
	"performance_manager/src/database"
	"performance_manager/src/entities"
)

type UserRepository struct {
	DB *database.DB
}

func (ur UserRepository) CreateUser(ctx context.Context, usr entities.CreateUser) (entities.User, error) {
	_, err := ur.DB.Exec(ctx,
		`
		INSERT INTO app_user (email, name, hashed_password)
		VALUES ($1, $2, $3);
	`,
		usr.Email, usr.Name, usr.Password,
	)
	if err != nil {
		return entities.User{}, nil
	}

	return entities.User{}, nil
}
