package transaction

import (
	"errors"
	"fmt"
	"store/modules/logins"
	"store/modules/products"

	"strconv"
	"time"
)

type Usecase struct {
	Repo        Repository
	ProductRepo products.Repository
	AdminRepo   logins.Repository
}

func (usecase Usecase) GetTransactions(limit int, offset int, colum string, sort string, search string) ([]Transaction, error) {
	Transactions, err := usecase.Repo.GetTransactions(limit, offset, colum, sort, search)
	return Transactions, err
}

func (usecase Usecase) GetTransaction(id string) (*Transaction, error) {
	transaction, err := usecase.Repo.GetTransaction(id)
	return transaction, err
}

func (usecase Usecase) CreateTransaction(req *Transaction) (*Transaction, error) {
	var transaction Transaction

	var total int
	for i, item := range req.Items {
		stringProduct := strconv.Itoa(item.ProductID)
		product, err := usecase.ProductRepo.GetProduct(stringProduct)
		if err != nil {
			return nil, err
		}
		if item.Quantity > product.Stock {
			return nil, errors.New("stock not enough")
		}
		product.Stock -= item.Quantity
		
		product.Sold += item.Quantity
		total += item.Quantity * product.Price
		req.Items[i].Price = product.Price
		err = usecase.ProductRepo.UpdateProductStockSold(item.ProductID, product)
		if err != nil {
			return nil, err
		}
	}
	transaction.AdminID = req.AdminID
	fmt.Println(req.BranchID)
	fmt.Println(transaction.BranchID)
	transaction.BranchID = req.BranchID
	fmt.Println(req.BranchID)
	fmt.Println(transaction.BranchID)
	transaction.Timestamp = time.Now()
	transaction.Total = total
	transaction.Items = req.Items
	fmt.Println("isi :", transaction.Items)

	err := usecase.Repo.CreateTransaction(&transaction)
	return &transaction, err
}
