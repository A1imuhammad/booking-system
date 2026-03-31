package jwt

import (
	"booking-system/services/auth/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user models.User, duration time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
