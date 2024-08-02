package cache

import (
	"context"
	"project-microservices/pkg/logger"
	"project-microservices/user_service/config"
	userServiceProto "project-microservices/user_service/proto"

	"github.com/go-redis/redis/v8"
)

type userCache struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

type UserCache interface {
	// set phoneNumber and sms sending data  with 3 minutes expired time
	SetVerificationKey(ctx context.Context, req *userServiceProto.PhoneVerificationReq, verificationKey string)
	// get phoneNumber or error
	GetVerificationKey(ctx context.Context, req *userServiceProto.UserPhoneCreateReq) (string, error)
}

func NewUserCache(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) UserCache {
	return &userCache{redisClient: redisClient, log: log, cfg: cfg}
}
