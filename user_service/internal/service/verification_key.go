package service

import (
	"context"
	userServiceProto "project-microservices/user_service/proto"

	uuid "github.com/satori/go.uuid"
)

func (u *userService) VerificationKey(ctx context.Context, req *userServiceProto.PhoneVerificationReq) (*userServiceProto.PhoneVerificationRes, error) {
	verificationKey := uuid.NewV4().String()
	u.cache.SetVerificationKey(ctx, req, verificationKey)

	return &userServiceProto.PhoneVerificationRes{VerificationKey: verificationKey}, nil
}
