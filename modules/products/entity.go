package products

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Respons struct {
	Message string    `json:"massage"`
	Data    []Product `json:"data"`
}
