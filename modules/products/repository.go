package products

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetProducts() ([]Product, error) {
	var products []Product
	result := repo.DB.Where("deleted_at is null").Find(&products)
	return products, result.Error
}

func (repo Repository) GetProduct(id string) (*Product, error) {
	var product *Product
	result := repo.DB.Where("deleted_at is null").First(&product, id)
	return product, result.Error
}

func (repo Repository) CreateProduct(product *Product) error {
	result := repo.DB.Select("Name", "Price", "Stock").Create(&product)
	return result.Error
}

func (repo Repository) UpdateProduct(id string, product *Product) (*Product, error) {
	result := repo.DB.Model(&Product{}).Where("deleted_at is null AND id = ?", id).Updates(&product)
	return product, result.Error
}

func (repo Repository) SoftDelete(id string, product *Product) (*Product, error) {
	result := repo.DB.Where("id = ?", id).Delete(&product)
	return product, result.Error
}
func (repo Repository) RestoreProduct(id string, product *Product) (*Product, error) {
	result := repo.DB.Unscoped().Model(&product).Where("id =? ", id).Update("deleted_at", gorm.DeletedAt{})
	return product, result.Error
}
