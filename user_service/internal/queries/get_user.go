package queries

import (
	"context"
	userServiceProto "project-microservices/user_service/proto"
)

func (u *userQueries) GetUser(ctx context.Context, req *userServiceProto.GetUser) (*userServiceProto.User, error) {
	return nil, nil
}
