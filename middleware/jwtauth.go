package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	commonHTTP "github.com/amine-bambrik-p8/go-lang-web-service/common/http"
	"github.com/dgrijalva/jwt-go"
)

// JWTAuthRequired wraps around a RouteHandler to check if user is logged in or not
// TODO should implement verification of expiration date
func JWTAuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "" {
			splitToken := strings.Split(token, "Bearer ")
			token = splitToken[1]
			token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Bad JWT Token")
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				commonHTTP.SendJSON(w, r, commonHTTP.ErrMessageUnauthorized, http.StatusUnauthorized)
				return
			}
			if token.Valid {
				next(w, r)
				return
			}
		}
		commonHTTP.SendJSON(w, r, commonHTTP.ErrMessageUnauthorized, http.StatusUnauthorized)
	}
}
