package service

import (
	"context"
	"project-microservices/pkg/middleware"
	userServiceProto "project-microservices/user_service/proto"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (u *userService) PhoneUserCreate(ctx context.Context, req *userServiceProto.UserPhoneCreateReq) (*userServiceProto.UserPhoneCreateRes, error) {
	phone, err := u.cache.GetVerificationKey(ctx, req)
	if err != nil {
		return nil, err
	}

	userId := uuid.NewV4().String()

	err = u.repo.CreateUserWithPhone(ctx, phone, userId)
	if err != nil {
		return nil, err
	}

	var userType middleware.UserType
	if req.UserType == 1 {
		userType = middleware.Admin
	} else {
		userType = middleware.User
	}

	refresh, err := u.middle.GenerateNewToken(middleware.GenerateTokenRequest{
		UserID:   userId,
		TokenFor: middleware.Refresh,
		UsedFor:  userType,
		ExpireAt: <-time.After(RefreshTokenExpired),
	}, "")
	if err != nil {
		return nil, err
	}
	access, err := u.middle.GenerateNewToken(middleware.GenerateTokenRequest{
		UserID:   userId,
		TokenFor: middleware.Access,
		UsedFor:  userType,
		ExpireAt: <-time.After(AccessTokenExpired),
	}, "")
	if err != nil {
		return nil, err
	}
	return &userServiceProto.UserPhoneCreateRes{
		AccessToken:  access.TokenString,
		RefreshToken: refresh.TokenString,
		UserType:     req.UserType,
	}, nil
}
