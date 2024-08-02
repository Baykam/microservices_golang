package repository

import (
	"context"
)

func (u *userRepository) CreateUserWithPhone(ctx context.Context, phone, userId string) error {

	_, err := u.sql.QueryContext(ctx, `INSERT INTO users(phone, userId) VALUES %s, %s`, phone, userId)
	if err != nil {
		return err
	}

	return nil
}
