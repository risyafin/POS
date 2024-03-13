package transaction

import (
	"store/modules/branches"
	"store/modules/logins"
	"store/modules/products"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        int                `json:"id"`
	Timestamp time.Time          `json:"timestamp"`
	Total     int                `json:"total"`
	BranchID  int                `json:"branch_id"`
	Branch    branches.Branch    `json:"branch"`
	AdminID   int                `json:"admin_id"`
	Admin     logins.Admin       `json:"admin"`
	Items     []TransactionsItem `json:"items"`
}

type TransactionItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type TransactionRequest struct {
	Items []TransactionItemRequest `json:"items"`
}

type TransactionsItem struct {
	gorm.Model
	ID            int              `json:"id"`
	ProductID     int              `json:"product_id"`
	TransactionID int              `json:"transaction_id"`
	Quantity      int              `json:"quantity"`
	Price         int              `json:"price"`
	Product       products.Product `json:"product"`
}

type GetAllTransactionsResponse struct {
	Message string
	Data    []Transaction
}

type GetDeteailTransactionResponse struct {
	Message string
	Data    *Transaction
}
