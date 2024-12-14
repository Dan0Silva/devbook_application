package authentication

import (
	"devbook_backend/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var mySigningKey = []byte(config.SecretKey)

func CreateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userID.String(),
		"exp":        time.Now().Add(time.Minute * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(mySigningKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetUserIDFromToken(r *http.Request) (string, error) {
	userTokenString := extractToken(r)
	userToken, err := jwt.Parse(userTokenString, getVerificationKey)

	if err != nil {
		return "", err
	}

	var userId interface{}
	if claims, ok := userToken.Claims.(jwt.MapClaims); !ok {
		return "", errors.New("map string error")
	} else {
		userId = claims["userId"]
	}

	userIdString := fmt.Sprintf("%v", userId)
	return userIdString, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if splitedToken := strings.Split(token, " "); len(splitedToken) == 2 {
		return splitedToken[1]
	}

	return ""
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
	}

	return mySigningKey, nil
}
