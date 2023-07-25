package branches

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetBrands(limit int, offset int, colum string, sort string, search string) ([]Branch, error) {
	var brands []Branch
	result := repo.DB.Limit(limit).Offset(offset).Order(colum+" "+sort).Where("name LIKE ?", "%"+search+"%").Find(&brands)
	return brands, result.Error
}
