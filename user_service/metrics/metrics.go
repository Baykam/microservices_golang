package metrics

import (
	"fmt"
	"project-microservices/user_service/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type UserServiceMetrics struct {
	SuccessGrpcSendData prometheus.Counter
	ErrorGrpcData       prometheus.Counter
	UserCreate          prometheus.Counter
	VerificationKey     prometheus.Counter
	UserUpdate          prometheus.Counter
	UserDelete          prometheus.Counter
}

func NewUserServiceMetrics(cfg *config.Config) *UserServiceMetrics {
	return &UserServiceMetrics{
		SuccessGrpcSendData: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorGrpcData: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		UserCreate: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		VerificationKey: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		UserUpdate: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		UserDelete: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
	}
}
