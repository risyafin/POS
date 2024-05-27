package logins

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Registration(admin *Admin) error {
	result := repo.DB.Select("Name", "Username", "Password", "BranchID").Create(&admin)
	return result.Error
}

func (repo Repository) Login(username string, password string) (*Admin, error) {
	var admin Admin
	result := repo.DB.Model(&admin).Where("username =? AND password =? ", username, password).First(&admin)
	return &admin, result.Error
}

func (repo Repository) GetAdmiByUsername(username string) (*Admin, error) {
	var admin *Admin
	result := repo.DB.Where("username = ?", username).First(&admin)
	return admin, result.Error
}
