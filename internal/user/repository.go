package user

import "webapi/internal/db"

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{Queries: queries}
}
