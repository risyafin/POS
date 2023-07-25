package logins

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
)

type Usecase struct {
	Repo Repository
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}
func CekPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

func (usecase Usecase) Registration(admin *Admin) error {
	hash, _ := HashPassword(admin.Password)
	admin.Password = hash
	err := usecase.Repo.Registration(admin)
	return err

}

func (usecase Usecase) Login(username, password string) (string, error) {
	admin, err := usecase.Repo.GetAdmiById(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("invalid idential")
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", errors.New("Invalid credentials")
	}

	claims := MyClaims{
		Username: admin.Username,
		Id:       admin.Id,
		BranchID: admin.BranchID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err1 := token.SignedString([]byte("Bolong"))
	if err != nil {
		return "", err1
	}

	return signedToken, nil
}
