package logins

import "github.com/golang-jwt/jwt"

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Id       int    `json:"id"`
}

