package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var mySigningKey = []byte("segredo")

func CreateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userID.String(),
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return signedToken, err
}
