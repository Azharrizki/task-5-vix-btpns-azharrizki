package helpers

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("ljsfdlkj9834ljksdn2355")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}