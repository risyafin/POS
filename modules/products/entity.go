package products

import (
	"errors"
	"time"
)

type Product struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Price     int        `json:"price"`
	Stock     int        `json:"stock"`
	Sold      int        `json:"sold"`
	DeletedAt *time.Time `json:"deleted_at"`
}
type Request struct {
	Status string `json:"status"`
}

type Respons struct {
	Message string    `json:"massage"`
	Data    []Product `json:"data"`
}

var (
	ErrProductAlreadyDeleted = errors.New("product already deleted")
	ErrProductNotDeleted     = errors.New("product is not deleted yet")
	ErrInvalidStatus         = errors.New("invalid status")
	ErrChangedStatus         = errors.New("status data cannot changed")
	ErrPoductHasBeenRemoved  = errors.New("product has been removed")
	ErrProductIdNotFound     = errors.New("data has been deleted")
)
