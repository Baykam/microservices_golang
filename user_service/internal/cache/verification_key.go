package cache

import (
	"context"
	"errors"
	userServiceProto "project-microservices/user_service/proto"
)

func (u *userCache) SetVerificationKey(ctx context.Context, req *userServiceProto.PhoneVerificationReq, verificationKey string) {
	phoneKey, smsKey := u.redisForVerificationKey(verificationKey)
	u.redisClient.Set(ctx, phoneKey, req.Phone, ExpiredTime)
	u.redisClient.Set(ctx, smsKey, "123456", ExpiredTime)
}

func (u *userCache) GetVerificationKey(ctx context.Context, req *userServiceProto.UserPhoneCreateReq) (string, error) {
	phoneKey, smsKey := u.redisForVerificationKey(req.VerificationKey)

	sms := u.redisClient.Get(ctx, smsKey)
	if sms.Val() != req.Sms {
		u.log.WarnMsg("sms value not true", errors.ErrUnsupported)
		return "", errors.New("sms value not true")
	}
	phone := u.redisClient.Get(ctx, phoneKey).Val()
	return phone, nil
}
