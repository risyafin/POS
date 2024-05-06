package products

import (
	"fmt"

	"time"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetProducts(limit int, offset int, colum string, sort string, keywords string) ([]Product, error) {
	products, err := usecase.Repo.GetProducts(limit, offset, colum, sort, keywords)
	return products, err
}

func (usecase Usecase) SearchingProduct(keywords string) ([]Product, error) {
	product, err := usecase.Repo.SearchingProduct(keywords)
	return product, err
}

func (usecase Usecase) Shorting(colum string) ([]Product, error) {
	product, err := usecase.Repo.Shorting(colum)
	return product, err
}

func (usecase Usecase) GetProduct(id string) (*Product, error) {
	var product *Product
	product, err := usecase.Repo.GetProduct(id)
	return product, err
}

func (usecase Usecase) UpdateProduct(id int, product *Product) error {
	product.ID = id
	if product.DeletedAt == nil {
		if err := usecase.Repo.UpdateProduct(id, product); err != nil {
			return ErrChangedStatus
		}
	}
	if product.DeletedAt != nil {
		return ErrPoductHasBeenRemoved
	}
	return nil
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

func (usecase Usecase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}
