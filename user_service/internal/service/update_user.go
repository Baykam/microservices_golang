package service

import (
	"context"
	userServiceProto "project-microservices/user_service/proto"
)

func (u *userService) UpdateUserData(ctx context.Context, req *userServiceProto.PostUser) (*userServiceProto.User, error) {
	return nil, nil
}
