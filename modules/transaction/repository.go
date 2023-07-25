package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetTransactions(limit int, offset int, colum string, sort string, search string) ([]Transaction, error) {
	var (
		transactions []Transaction
		db           = repo.DB
	)
	db = db.Preload("Branch")
	db = db.Preload("Admin")
	result := db.Limit(limit).Offset(offset).
		Order(colum+" "+sort).Where("id LIKE ?", "%"+search+"%").Find(&transactions)
	return transactions, result.Error
}

func (repo Repository) GetTransaction(id string) (*Transaction, error) {
	var transaction *Transaction
	result := repo.DB.Preload("Branch").Preload("Admin").Preload("Items.Product").First(&transaction, id)
	return transaction, result.Error
}

func (repo Repository) CreateTransaction(transaction *Transaction) error {
	result := repo.DB.Create(&transaction)
	return result.Error
}
