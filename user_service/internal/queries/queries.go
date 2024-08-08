package queries

import (
	"context"
	"project-microservices/pkg/middleware"
	"project-microservices/user_service/config"
	"project-microservices/user_service/internal/cache"
	"project-microservices/user_service/internal/repository"
	userServiceProto "project-microservices/user_service/proto"
)

type UserQueries interface {
	VerificationKey(ctx context.Context, req *userServiceProto.PhoneVerificationReq) (*userServiceProto.PhoneVerificationRes, error)
	PhoneUserCreate(ctx context.Context, req *userServiceProto.UserPhoneCreateReq) (*userServiceProto.UserPhoneCreateRes, error)
	GetUser(ctx context.Context, req *userServiceProto.GetUser) (*userServiceProto.User, error)
	UpdateUserData(ctx context.Context, req *userServiceProto.PostUser) (*userServiceProto.User, error)
}

type userQueries struct {
	repo   repository.UserRepository
	cache  cache.UserCache
	middle middleware.MiddlewareAuth
	cfg    *config.Config
}

func NewUserQueries(repo repository.UserRepository,
	cache cache.UserCache,
	cfg *config.Config) UserQueries {
	return &userQueries{repo: repo, cache: cache, cfg: cfg}
}
