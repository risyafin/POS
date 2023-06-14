package logins

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) Registration (admin *Admin) error{
	err := usecase.Repo.Registration(admin)
	return err
}

func (usecase Usecase) Login(username string, password string) (string, error) {
	admin, err := usecase.Repo.Login(username, password)
	if err != nil {
		err.Error()
	}
	claims := MyClaims{
		Username: admin.Username,
		Id:       admin.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Bolong"))

	return signedToken, err
}
