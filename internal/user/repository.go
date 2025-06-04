package user

import (
	"context"
	"webapi/internal/db"
)

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{Queries: queries}
}

func (r *UserRepository) StoreUser(ctx context.Context, input *UserInput) (*db.User, error) {
	user, err := r.Queries.CreateUser(ctx, db.CreateUserParams{
		Username:  input.Username,
		Password:  input.Password,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
