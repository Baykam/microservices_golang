package userHttp

import (
	"net/http"
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/internal/user/service"
	"project-microservices/api_gateway_service/metrics"
	"project-microservices/dto"
	"project-microservices/pkg/constants"
	"project-microservices/pkg/logger"
	"project-microservices/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type userHandlers struct {
	cfg     *config.Config
	log     logger.Logger
	v       validator.Validate
	engine  *gin.Engine
	metrics metrics.UserMetrics
	us      *service.UserService
	middle  middleware.MiddlewareAuth
}

func NewUserHandlers(
	cfg *config.Config,
	log logger.Logger,
	v validator.Validate,
	engine *gin.Engine,
	metrics metrics.UserMetrics,
	us *service.UserService,
) *userHandlers {
	return &userHandlers{cfg: cfg, log: log, v: v, engine: engine, metrics: metrics, us: us}
}

// LoginUser For phone
// @Tags Auth
// @Summary loginUser
// @Description Phone needed verificationKey && SMS
// @Accept json
// @Produce json
// @Param   verificationKey body dto.UserCreateReq true "Verification Key And SMS Request"
// @Success 201 {object} dto.UserCreateRes
// @Router /auth/login [post]
func (u *userHandlers) LoginUser(c *gin.Context) {
	u.metrics.UserCreateRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(c, "apiGateway.userHandler.VerificationKey")

	defer span.Finish()

	req := dto.UserCreateReq{}
	if err := c.BindJSON(req); err != nil {
		u.log.WarnMsg("BindJson", err)
		u.traceError(span, err)
		return
	}

	res, err := u.us.Queries.CreateUserWithPhone(ctx, req)
	if err != nil {
		u.log.WarnMsg("CreateUserWithPhone", err)
		u.traceError(span, err)
		return
	}

	u.metrics.UserCreateRequests.Inc()
	c.JSON(http.StatusCreated, res)
}

// VerificationKey For phone
// @Tags Auth
// @Summary verification for Phone
// @Description Phone send
// @Accept json
// @Produce json
// @Param   phone_verification body dto.PhoneVerificationReq true "Phone Verification Request"
// @Success 200 {object} dto.PhoneVerificationRes
// @Router /auth/verification [post]
func (u *userHandlers) VerificationKey(c *gin.Context) {
	u.metrics.VerificationKeyRequests.Inc()

	sp, ctx := opentracing.StartSpanFromContext(c, "apiGateway.userHandler.VerificationKey")
	defer sp.Finish()

	req := &dto.PhoneVerificationReq{}
	if err := c.BindJSON(req); err != nil {
		u.log.WarnMsg("BindJson", err)
		u.traceError(sp, err)
		return
	}

	res, err := u.us.Queries.GetVerificationKey(ctx, req)
	if err != nil {
		u.log.WarnMsg("GetVerificationKey", err)
		u.traceError(sp, err)
		return
	}

	u.metrics.SuccessHttpRequests.Inc()
	c.JSON(http.StatusOK, res)
}

// Update User
// @Tags Auth
// @Summary Update User
// @Description Update User
// @Accept json
// @Produce json
// @Param   user_update body dto.UserUpdateReq true "Update User Request"
// @Success 200 {object} dto.User
// @Router /auth [post]
func (u *userHandlers) UpdateUser(c *gin.Context) {
	u.metrics.UserUpdateRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(c, "apiGateway.userHandler.UpdateUser")
	defer span.Finish()
	access := c.GetHeader(constants.AccessToken)

	userId, err := u.middle.ValidateToken(middleware.VerifyTokenRequest{
		TokenString: access,
		UsedFor:     middleware.User,
		TokenFor:    middleware.Access,
	}, u.cfg.ServiceName)
	if err != nil {
		u.log.WarnMsg("Validate Token", err)
		u.traceError(span, err)
		return
	}

	req := &dto.UserUpdateReq{}
	if err := c.BindJSON(req); err != nil {
		u.log.WarnMsg("BindJson", err)
		u.traceError(span, err)
		return
	}
	req.UserId = userId.UserID

	err = u.us.Commands.UpdateUser(ctx, req)

	if err != nil {
		u.log.WarnMsg("UpdateUser", err)
		u.traceError(span, err)
		return
	}

	res := &dto.User{
		Phone:     req.Phone,
		Email:     &req.Email,
		UpdatedAt: time.Now(),
	}

	u.metrics.SuccessHttpRequests.Inc()
	c.JSON(http.StatusOK, res)
}

// Get User
// @Tags Auth
// @Summary Get User
// @Description Get User
// @Accept json
// @Produce json
// @Success 200 {object} dto.User
// @Router /auth [get]
func (u *userHandlers) GetUser(c *gin.Context) {
	u.metrics.UserGetRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(c, "apiGateway.userHandler.GetUser")
	defer span.Finish()

	access := c.GetHeader(constants.AccessToken)

	userId, err := u.middle.ValidateToken(middleware.VerifyTokenRequest{
		TokenString: access,
		UsedFor:     middleware.User,
		TokenFor:    middleware.Access,
	}, u.cfg.ServiceName)
	if err != nil {
		return
	}

	res, err := u.us.Queries.GetUser(ctx, dto.GetUserReq{Token: userId.UserID})
	if err != nil {
		return
	}

	u.metrics.SuccessHttpRequests.Inc()
	c.JSON(http.StatusOK, res)
}
