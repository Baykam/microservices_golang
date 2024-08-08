package queries

import (
	"context"
	userServiceProto "project-microservices/user_service/proto"
)

func (u *userQueries) UpdateUserData(ctx context.Context, req *userServiceProto.PostUser) (*userServiceProto.User, error) {
	return nil, nil
}
