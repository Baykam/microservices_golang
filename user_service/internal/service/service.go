package service

import (
	"project-microservices/user_service/config"
	"project-microservices/user_service/internal/cache"
	"project-microservices/user_service/internal/commands"
	"project-microservices/user_service/internal/queries"
	"project-microservices/user_service/internal/repository"
)

type UserService struct {
	Queries  queries.UserQueries
	Commands commands.Commands
}

func NewUserService(Repo repository.UserRepository, Cache cache.UserCache, cfg *config.Config) *UserService {
	queries := queries.NewUserQueries(Repo, Cache, cfg)
	commands := commands.NewUserCommands(Repo, Cache, cfg)
	return &UserService{Queries: queries, Commands: commands}
}
