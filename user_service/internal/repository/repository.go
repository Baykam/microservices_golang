package repository

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	sql   *sql.DB
	mongo mongo.Client
}

type UserRepository interface {
	CreateUserWithPhone(ctx context.Context, phone, userId string) error
}

func NewUserRepository(sql *sql.DB,
	mongo mongo.Client,
) UserRepository {
	return &userRepository{sql: sql, mongo: mongo}
}
