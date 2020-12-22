package security

import (
	"backend-github-trending/model"

	"github.com/dgrijalva/jwt-go"
)

// GenToken for generate token for user to use other api
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserID:         user.UserID,
		Role:           user.Role,
		StandardClaims: jwt.StandardClaims{},
	}

	return "nil", nil
}
