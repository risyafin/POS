package products

import (
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	Status    string    `json:"status"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Respons struct {
	Message string    `json:"massage"`
	Data    []Product `json:"data"`
}
