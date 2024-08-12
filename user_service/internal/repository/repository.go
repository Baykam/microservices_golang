package repository

import (
	"context"
	"database/sql"
)

type userRepository struct {
	sql *sql.DB
	// mongo mongo.Client
}

type UserRepository interface {
	CreateUserWithPhone(ctx context.Context, phone, userId string) error
	GetUser()
}

func NewUserRepository(sql *sql.DB) UserRepository {
	return &userRepository{sql: sql}
}
