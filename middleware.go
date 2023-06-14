package main

import (
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
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		_, ok := token.Claims.(*logins.MyClaims) //forward ke usecase create transaction
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token invalid"))
			return
		}
		// contextClaims := context.Background()
		// // contextClaimsUser := context.Background()

		// adminId := context.WithValue(contextClaims, "adminId", claims.Id)
		// username := context.WithValue(adminId, "username", claims.Username)
		// // contextClaimsUser = context.WithValue(contextClaimsUser, "adminUser", claims.Username)
		next(w, r)

	}
}
