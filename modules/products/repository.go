package products

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetProducts(limit int, offset int, colum string, sort string, keywords string) ([]Product, error) {
	var products []Product
	result := repo.DB.Limit(limit).Offset(offset).Order(colum+" "+sort).Where("name LIKE ?", "%"+keywords+"%").Find(&products)
	return products, result.Error
}

func (repo Repository) SearchingProduct(keywords string) ([]Product, error) {
	var product []Product
	result := repo.DB.Where("name LIKE ?", "%"+keywords+"%").Find(&product)
	return product, result.Error
}

func (repo Repository) Shorting(colum string) ([]Product, error) {
	var product []Product
	result := repo.DB.Order(colum).Find(&product)
	return product, result.Error
}

// func (repo Repository) GetProducts() ([]Product, error) {
// 	var products []Product
// 	result := repo.DB.Unscoped().Find(&products)
// 	return products, result.Error
// }

func (repo Repository) GetProduct(id string) (*Product, error) {
	var product *Product
	result := repo.DB.First(&product, id)
	return product, result.Error
}

func (repo Repository) CreateProduct(product *Product) error {
	result := repo.DB.Select("Name", "Price", "Stock").Create(&product)
	return result.Error
}

func (repo Repository) UpdateProduct(id int, product *Product) error {
	result := repo.DB.Model(&Product{}).Select("Name", "Price", "Stock").Where(id).Updates(&product) // handle di usecase
	return result.Error
}

func (repo Repository) UpdateProductStockSold(id int, product *Product) error {
	result := repo.DB.Model(&Product{}).Select("Stock", "Sold").Where(id).Updates(&product) // handle di usecase
	return result.Error
}

func (repo Repository) Save(product *Product) error {
	result := repo.DB.Save(&product)
	return result.Error
}

func (repo Repository) SoftDelete(id string, product *Product) (*Product, error) {
	result := repo.DB.Where("id = ?", id).Delete(&product)
	return product, result.Error
}
func (repo Repository) RestoreProduct(id string, product *Product) (*Product, error) {
	result := repo.DB.Unscoped().Model(&product).Where("id =? ", id).Update("deleted_at", gorm.DeletedAt{})
	return product, result.Error
}
