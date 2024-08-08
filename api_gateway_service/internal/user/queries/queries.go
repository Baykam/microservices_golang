package queries

import (
	"context"
	"project-microservices/dto"
	"project-microservices/mappers"
	userServiceProto "project-microservices/user_service/proto"
)

type UserQueries interface {
	GetVerificationKey(ctx context.Context, req *dto.PhoneVerificationReq) (*dto.PhoneVerificationRes, error)
	CreateUserWithPhone(ctx context.Context, req dto.UserCreateReq) (*dto.UserCreateRes, error)
	GetUser(ctx context.Context, req dto.GetUserReq) (*dto.User, error)
	UpdateUser(ctx context.Context, req dto.UserUpdateReq) (*dto.User, error)
}

type userQueries struct {
	client userServiceProto.UserServiceClient
}

func NewUserQueries(client userServiceProto.UserServiceClient) UserQueries {
	return &userQueries{client: client}
}

func (q *userQueries) CreateUserWithPhone(ctx context.Context, req dto.UserCreateReq) (*dto.UserCreateRes, error) {
	res, err := q.client.PhoneUserCreate(ctx, mappers.UserPhoneCreateToGrpc(req))
	if err != nil {
		return nil, err
	}
	createDto := mappers.UserPhoneCreateFromGrpc(res)
	return createDto, nil
}

func (q *userQueries) GetVerificationKey(ctx context.Context, req *dto.PhoneVerificationReq) (*dto.PhoneVerificationRes, error) {
	res, err := q.client.VerificationKey(ctx, mappers.UserVerificationKeyToGrpc(req))
	if err != nil {
		return nil, err
	}
	createDto := mappers.UserVerificationKeyFromGrpc(res)
	return createDto, nil
}

func (q *userQueries) GetUser(ctx context.Context, req dto.GetUserReq) (*dto.User, error) {
	res, err := q.client.GetUser(ctx, mappers.GetUserToGrpc(req))
	if err != nil {
		return nil, err
	}
	last := mappers.GetUserFromGrpc(res)
	return last, nil
}

func (q *userQueries) UpdateUser(ctx context.Context, req dto.UserUpdateReq) (*dto.User, error) {
	res, err := q.client.UpdateUserData(ctx, mappers.UserUpdateToGrpc(req))
	if err != nil {
		return nil, err
	}
	last := mappers.GetUserFromGrpc(res)
	return last, nil
}
