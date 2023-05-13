package transaction

import (
	"store/modules/products"
	"time"
)

type Transaction struct {
	ID        int                `json:"id"`
	Timestamp time.Time          `json:"timestamp"`
	Total     int                `json:"total"`
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
