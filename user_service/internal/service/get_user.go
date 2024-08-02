package service

import (
	"context"
	userServiceProto "project-microservices/user_service/proto"
)

func (u *userService) GetUser(ctx context.Context, req *userServiceProto.GetUser) (*userServiceProto.User, error) {
	return nil, nil
}
