package mongoRepository

import (
	"project-microservices/pkg/logger"
	"project-microservices/user_service/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepo struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepo(log logger.Logger, cfg *config.Config, db *mongo.Client) *userMongoRepo {
	return &userMongoRepo{log: log, cfg: cfg, db: db}
}

func (u *userMongoRepo) CreateUser() {

}
