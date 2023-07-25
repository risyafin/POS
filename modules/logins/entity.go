package logins

import "github.com/golang-jwt/jwt"

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	BranchID int    `json:"branch_id"`
}
type AdminResponses struct {
	Username string `json:"username"`
	BranchID int    `json:"branch_id"`
}

func (admin AdminResponses) TableName() string {
	return "admins"
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Id       int    `json:"id"`
	BranchID int    `json:"branch_id"`
}
