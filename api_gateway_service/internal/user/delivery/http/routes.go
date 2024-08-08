package userHttp

func (u *userHandlers) Run() {
	// auth := u.engine.Group("/auth")
	auth := u.engine
	auth.POST("/verification", u.VerificationKey)
	auth.POST("/login", u.LoginUser)
	auth.POST("", u.UpdateUser)
	auth.GET("", u.GetUser)
}
