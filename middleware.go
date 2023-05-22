package main

import (
	"context"
	"net/http"
	"store/modules/logins"

	"github.com/golang-jwt/jwt"
)

func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &logins.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("Bolong"), nil
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		claims, ok := token.Claims.(*logins.MyClaims) //forward ke usecase create transaction
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token invalid"))
			return
		}
		contextClaims := context.Background()
		// contextClaimsUser := context.Background()

		contextClaims = context.WithValue(contextClaims, "adminId", claims.Id)
		// contextClaimsUser = context.WithValue(contextClaimsUser, "adminUser", claims.Username)

		r = r.WithContext(contextClaims)

	}
}
