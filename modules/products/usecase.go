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

func (usecase Usecase) UpdateProduct(id string, product *Product) (*Product,error) {
	product , err := usecase.Repo.UpdateProduct(id, product)
	return product, err

}

func (usecase Usecase) SoftDelete(id string, req *Request) (*Product, error) {
	var product Product
	var err error
	if req.Status == "inactive" {
		product, err := usecase.Repo.SoftDelete(id, &product)
		return product, err
	}
	if req.Status == "active" {
		product, err := usecase.Repo.RestoreProduct(id, &product)
		return product, err
	}
	return &product, err
}

// func (usecase Usecase) RestoreProduct(id string, product *Product) (*Product, error) {
// 	product, err := usecase.Repo.RestoreProduct(id, product)
// 	return product, err
// }

func (usecase Usecase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}
