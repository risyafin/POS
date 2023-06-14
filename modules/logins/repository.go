package logins

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Registration(admin *Admin) error {
	result := repo.DB.Select("Name", "Username", "Password").Create(&admin)
	return result.Error
}

func (repo Repository) Login(username string, password string) (*Admin, error) {
	var admin Admin
	result := repo.DB.Model(&admin).Where("username =? AND password =? ", username, password).First(&admin)
	fmt.Println(&admin)
	return &admin, result.Error
}

func (repo Repository) GetAdmiById(id int) (*Admin, error) {
	var admin *Admin
	result := repo.DB.First(&admin, id)
	return admin, result.Error
}
