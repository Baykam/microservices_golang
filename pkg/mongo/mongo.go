package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeOut  = 30 * time.Second
	macConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

type Config struct {
	Uri      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func NewMongoDbConn(ctx context.Context, cfg *Config) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(cfg.Uri).
		SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password}).
		SetConnectTimeout(connectTimeOut).
		SetMaxConnIdleTime(macConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	fmt.Println(clientOptions)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
