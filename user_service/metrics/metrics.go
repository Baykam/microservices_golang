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
			Name: fmt.Sprintf("%s_success_grpc_send_data_total", cfg.ServiceName),
			Help: "The total number of successful gRPC data sends",
		}),
		ErrorGrpcData: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_data_total", cfg.ServiceName),
			Help: "The total number of gRPC data errors",
		}),
		UserCreate: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_create_total", cfg.ServiceName),
			Help: "The total number of users created",
		}),
		VerificationKey: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_verification_key_total", cfg.ServiceName),
			Help: "The total number of verification keys generated",
		}),
		UserUpdate: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_update_total", cfg.ServiceName),
			Help: "The total number of users updated",
		}),
		UserDelete: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_delete_total", cfg.ServiceName),
			Help: "The total number of users deleted",
		}),
	}
}
