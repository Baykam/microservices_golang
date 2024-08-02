package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidUserType    = errors.New("invalid user type")
	ErrInvalidToken       = errors.New("invalid token")
	ErrFailedToParseToken = errors.New("failed to parse token to claims")
	ErrExpiredToken       = errors.New("token expired")
	ErrInvalidTokenType   = errors.New("token type not accepted access or refresh")
)

type jwtClaims struct {
	TokenType TokenType
	UserId    string
	UserType  UserType
	ExpiresAt time.Time
}

func (c *jwtClaims) Valid() error {
	if time.Since(c.ExpiresAt) > 0 {
		return ErrExpiredToken
	}
	return nil
}

func (c *jwtClaims) GenerateNewToken(req GenerateTokenRequest, secretKey string) (*GenerateTokenResponse, error) {
	newsecretKey := []byte(secretKey)

	claims := &jwtClaims{
		TokenType: req.TokenFor,
		UserType:  req.UsedFor,
		UserId:    req.UserID,
		ExpiresAt: req.ExpireAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(newsecretKey)
	if err != nil {
		return nil, err
	}
	response := &GenerateTokenResponse{
		TokenString: tokenString,
	}
	return response, nil
}

func (c *jwtClaims) ValidateToken(req VerifyTokenRequest, secretKey string) (*VerifyTokenResponse, error) {
	secretkey := []byte(secretKey)

	if req.UsedFor != Admin && req.UsedFor != User {
		return &VerifyTokenResponse{}, ErrInvalidUserType
	}

	if req.TokenFor != Access && req.TokenFor != Refresh {
		return &VerifyTokenResponse{}, ErrInvalidTokenType
	}

	token, err := jwt.ParseWithClaims(req.TokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return &VerifyTokenResponse{}, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return &VerifyTokenResponse{}, ErrFailedToParseToken
	}

	if err := claims.Valid(); err != nil {
		return &VerifyTokenResponse{}, err
	}

	response := &VerifyTokenResponse{
		UserID: claims.UserId,
	}

	return response, nil

}
