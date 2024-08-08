package commands

import (
	"context"
	"project-microservices/pkg/middleware"
	"project-microservices/user_service/config"
	"project-microservices/user_service/internal/cache"
	"project-microservices/user_service/internal/repository"
)

type Commands interface {
	UpdateUser(ctx context.Context) error
}

type commands struct {
	repo   repository.UserRepository
	cache  cache.UserCache
	middle middleware.MiddlewareAuth
	cfg    *config.Config
}

func NewUserCommands(repo repository.UserRepository,
	cache cache.UserCache,
	cfg *config.Config) Commands {
	return &commands{repo: repo, cache: cache, cfg: cfg}
}
