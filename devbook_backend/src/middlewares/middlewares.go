package middlewares

import (
	"devbook_backend/src/authentication"
	"devbook_backend/src/response"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.Error(w, "Not logged in", http.StatusUnauthorized, err.Error())
			return
		}

		next(w, r)
	}
}
