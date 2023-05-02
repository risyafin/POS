package products

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Respon struct {
	Massage string    `json:"massage"`
	Data    []Product `json:"data"`
}
