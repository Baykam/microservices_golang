package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
	Password string `yaml:"password"`
}

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func ConnPostgres(cfg *Config) (*sql.DB, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxIdleTime(minConns)
	db.SetConnMaxIdleTime(maxConnIdleTime)
	db.SetConnMaxLifetime(maxConnLifetime)

	go func() {
		for {
			if err := db.PingContext(ctx); err != nil {
				fmt.Printf("PostgreSQL health check failed: %v\n", err)
			}
			time.Sleep(healthCheckPeriod)
		}
	}()

	return db, nil
}
