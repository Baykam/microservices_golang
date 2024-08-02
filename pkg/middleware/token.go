package middleware

import (
	"time"
)

type UserType string
type TokenType string

type MiddlewareAuth interface {
	GenerateNewToken(req GenerateTokenRequest, secretKey string) (*GenerateTokenResponse, error)
	ValidateToken(req VerifyTokenRequest, secretKey string) (*VerifyTokenResponse, error)
}

const (
	Refresh TokenType = "refresh"
	Access  TokenType = "access"
)

const (
	Admin UserType = "admin"
	User  UserType = "user"
)

type GenerateTokenRequest struct {
	UserID   string
	TokenFor TokenType
	UsedFor  UserType
	ExpireAt time.Time
}

type GenerateTokenResponse struct {
	TokenString string
}

type VerifyTokenRequest struct {
	TokenString string
	UsedFor     UserType
	TokenFor    TokenType
}

type VerifyTokenResponse struct {
	UserID string
}
