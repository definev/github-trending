package model

import "github.com/dgrijalva/jwt-go"

// JwtCustomClaims struct
type JwtCustomClaims struct {
	UserID         string
	Role           string
	StandardClaims jwt.StandardClaims
}
