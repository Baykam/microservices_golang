package metrics

import (
	"fmt"
	"project-microservices/api_gateway_service/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type UserMetrics struct {
	SuccessHttpRequests     prometheus.Counter
	ErrorHttpRequests       prometheus.Counter
	UserCreateRequests      prometheus.Counter
	UserUpdateRequests      prometheus.Counter
	VerificationKeyRequests prometheus.Counter
	UserGetRequests         prometheus.Counter
}

func NewUserMetrics(cfg *config.Config) *UserMetrics {
	return &UserMetrics{
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_request_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requests_total", cfg.ServiceName),
			Help: "The total number of error http requests",
		}),
		UserCreateRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_create_requests_total", cfg.ServiceName),
			Help: "The total number of user create requests",
		}),
		UserUpdateRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_update_requests_total", cfg.ServiceName),
			Help: "The total number of user update requests",
		}),
		VerificationKeyRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_verification_key_requests_total", cfg.ServiceName),
			Help: "The total number of verification key requests",
		}),
		UserGetRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_user_requests_total", cfg.ServiceName),
			Help: "The total number of get user requests",
		}),
	}
}
