package store

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetBrands(limit int, offset int, colum string, sort string, search string) ([]Brand, error) {
	var brands []Brand
	result := repo.DB.Limit(limit).Offset(offset).Order(colum+" "+sort).Where("shop LIKE ?", "%"+search+"%").Find(&brands)
	return brands, result.Error
}
