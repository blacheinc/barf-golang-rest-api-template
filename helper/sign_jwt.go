package helper

import (
	"time"

	"github.com/blacheinc/pixel/primer"
	"github.com/blacheinc/pixel/types"
	"github.com/golang-jwt/jwt/v4"
)

// SignJWT signs a JWT with the given address
func SignJWT(id string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, types.JWTClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			// expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * 30 * time.Hour)),
			Issuer:    primer.ENV.AppName,
		},
	}).SignedString([]byte(primer.ENV.JWTSecret))
}
