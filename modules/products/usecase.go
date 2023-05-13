package products

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetProducts() ([]Product, error) {
	products, err := usecase.Repo.GetProducts()
	return products, err
}

func (usecase Usecase) GetProduct(id string) (*Product, error) {
	product, err := usecase.Repo.GetProduct(id)
	return product, err
}

func (usecase Usecase) UpdateProduct(id string, product *Product) error {
	err := usecase.Repo.UpdateProduct(id, product)
	return err

}

func (usecase Usecase) SoftDelete(id string, product *Product) (*Product, error) {
	product, err := usecase.Repo.SoftDelete(id, product)
	return product, err
}

// func (usecase Usecase) RestoreProduct(id string, product *Product) (*Product, error) {
// 	product, err := usecase.Repo.RestoreProduct(id, product)
// 	return product, err
// }

func (usecase Usecase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}
