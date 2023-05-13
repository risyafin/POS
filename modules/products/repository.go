package products

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetProducts() ([]Product, error) {
	var products []Product
	result := repo.DB.Find(&products, "status = ? AND deleted_at is null", "active")
	return products, result.Error
}

func (repo Repository) GetProduct(id string) (*Product, error) {
	var product *Product
	result := repo.DB.First(&product, id, "status = ? AND deleted_at is null", "active")
	return product, result.Error
}

func (repo Repository) CreateProduct(product *Product) error {
	result := repo.DB.Create(&product)
	return result.Error
}

func (repo Repository) UpdateProduct(id string, product *Product) error {
	result := repo.DB.Model(&Product{}).Where("id = ?", id, "status = ? AND deleted_at is null", "active").Updates(&product)
	return result.Error
}

func (repo Repository) SoftDelete(id string, product *Product) (*Product, error) {
	result := repo.DB.Model(&Product{}).Where("id =?", id).Updates(&product)
	return product, result.Error
}

func (repo Repository) RestoreProduct(id string, product *Product) (*Product, error) {
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates("deleted_at is null AND status = active")
	return product, result.Error
}
