package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID     uuid.UUID `json:"id"`
	Phone  string    `json:"phone"`
	UserId string    `json:"user_id"`
	Email  *string   `json:"email"`
}

type PhoneVerificationReq struct {
	Phone string `json:"phone"`
}

type PhoneVerificationRes struct {
	VerificationKey string `json:"verification_key"`
}

type UserCreateReq struct {
	SMS             string `json:"sms"`
	VerificationKey string `json:"verification_key"`
}

type UserCreateRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserType     int    `json:"user_type"`
}

type JWTRefreshTokenApiReq struct {
	UserType     int    `json:"user_type"`
	RefreshToken string `json:"refresh_token"`
}

type JWTRefreshTokenApiRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserType     int    `json:"user_type"`
}

type GetUserReq struct {
	Token string
}

type UserUpdateReq struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`

	UpdatedAt time.Time
	UserId    string
}
