package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}
