package products

import (
	"fmt"
	"time"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetProducts() ([]Product, error) {
	products, err := usecase.Repo.GetProducts()
	return products, err
}

func (usecase Usecase) GetProduct(id string) (*Product, error) {
	var product *Product
	if product.DeletedAt == nil {
		if product, err := usecase.Repo.GetProduct(id); err != nil {
			return product, ErrChangedStatus
		}

	}
	if product.DeletedAt != nil {
		return nil, ErrPoductHasBeenRemoved
	}
	return product, nil
}

func (usecase Usecase) UpdateProduct(id int, product *Product) error {
	err := usecase.Repo.UpdateProduct(id, product)
	product.ID = id
	return err

}

func (usecase Usecase) SoftDelete(id string, status string) (*Product, error) {
	product, err := usecase.Repo.GetProduct(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(product.DeletedAt)
	if status == "active" {
		fmt.Println("deleted=active", product.DeletedAt)
		if product.DeletedAt == nil {
			return nil, ErrProductNotDeleted
		} else if product.DeletedAt != nil {
			product.DeletedAt = nil
			fmt.Println(product.DeletedAt)
		}
	} else if status == "inactive" {
		fmt.Println("delete=inactive", product.DeletedAt)
		if product.DeletedAt == nil {
			fmt.Println(product.DeletedAt)
			deleteAt := time.Now()
			product.DeletedAt = &deleteAt
			fmt.Println(product.DeletedAt)
		} else if product.DeletedAt != nil {
			return nil, ErrProductAlreadyDeleted
		}
	} else {
		return nil, ErrInvalidStatus
	}
	if err := usecase.Repo.Save(product); err != nil {
		return nil, ErrChangedStatus
	}

	return product, nil

}

// func (usecase Usecase) RestoreProduct(id string, product *Product) (*Product, error) {
// 	product, err := usecase.Repo.RestoreProduct(id, product)
// 	return product, err
// }

func (usecase Usecase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}
