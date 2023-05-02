package products

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetProducts() ([]Product, error) {
	products, err := usecase.Repo.GetProducts()
	return products, err
}
