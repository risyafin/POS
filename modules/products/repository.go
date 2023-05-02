package products

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetProducts() ([]Product,error){
	var products []Product
	result :=repo.DB.Find(&products)
	return products, result.Error
}