package service

import (
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/internal/user/commands"
	"project-microservices/api_gateway_service/internal/user/queries"
	productKafka "project-microservices/pkg/kafka"
	"project-microservices/pkg/logger"
	userServiceProto "project-microservices/user_service/proto"
)

type UserService struct {
	Queries  queries.UserQueries
	Commands commands.Commands
}

func NewUserService(log logger.Logger, cfg *config.Config, kafkaProducers productKafka.Producer, userProto userServiceProto.UserServiceClient) *UserService {
	queries := queries.NewUserQueries(userProto)
	commands := commands.NewCommandHandler(cfg, kafkaProducers)
	return &UserService{Queries: queries, Commands: commands}
}
