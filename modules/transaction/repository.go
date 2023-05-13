package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetTransactions() ([]Transaction, error) {
	var transactions []Transaction
	result := repo.DB.Find(&transactions)
	return transactions, result.Error
}

func (repo Repository) GetTransaction(id string) (*Transaction, error) {
	var transaction *Transaction
	result := repo.DB.Preload("Items.Product").First(&transaction, id)
	return transaction, result.Error
}

func (repo Repository) CreateTransaction(transaction *Transaction) error {
	result := repo.DB.Create(&transaction)
	return result.Error
}
