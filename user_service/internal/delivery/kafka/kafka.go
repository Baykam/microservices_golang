package userKafkaConn

import (
	"context"
	"fmt"
	"project-microservices/pkg/logger"
	"project-microservices/user_service/config"
	"project-microservices/user_service/internal/service"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
)

const PoolSize int = 30

type userMessageProcessor struct {
	log logger.Logger
	cfg *config.Config
	v   *validator.Validate
	us  service.UserService
}

func NewUserMessagesProcessor(log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	us service.UserService) *userMessageProcessor {
	return &userMessageProcessor{log: log, cfg: cfg, v: v, us: us}
}

func (u *userMessageProcessor) ProcessMessages(ctx context.Context, reader *kafka.Reader, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		message, err := reader.FetchMessage(ctx)
		if err != nil {
			// u.log.Infof("workerId : %v, err : %v", workerId, err)
			fmt.Printf("workerId : %v, err : %v", workerId, err)
			continue
		}

		switch message.Topic {
		case u.cfg.KafkaTopics.UserDeleted.TopicName:
		case u.cfg.KafkaTopics.UserUpdated.TopicName:
		}
	}
}
