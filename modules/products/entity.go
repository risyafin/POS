package products

import "gorm.io/gorm"

type Product struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Price     int            `json:"price"`
	Stock     int            `json:"stock"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type Request struct {
	Status string  `json:"status"`
}

type Respons struct {
	Message string    `json:"massage"`
	Data    []Product `json:"data"`
}
