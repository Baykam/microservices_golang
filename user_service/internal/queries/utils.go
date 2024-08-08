package queries

import "time"

const (
	RefreshTokenExpired = time.Hour * 360
	AccessTokenExpired  = time.Minute * 15
)
