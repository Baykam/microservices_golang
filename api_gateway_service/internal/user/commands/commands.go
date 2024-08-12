package commands

import (
	"context"
	"project-microservices/api_gateway_service/config"
	"project-microservices/dto"
	productKafka "project-microservices/pkg/kafka"
)

type Commands interface {
	UpdateUser(ctx context.Context, req *dto.UserUpdateReq) error
}

type commandsHandler struct {
	cfg           *config.Config
	kafkaProducer productKafka.Producer
}

func NewCommandHandler(
	cfg *config.Config,
	kafkaProducer productKafka.Producer) Commands {
	return &commandsHandler{cfg: cfg, kafkaProducer: kafkaProducer}
}
