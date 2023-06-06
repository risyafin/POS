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

func (usecase Usecase) GetTransactions() ([]Transaction, error) {
	Transactions, err := usecase.Repo.GetTransactions()
	return Transactions, err
}

func (usecase Usecase) GetTransaction(id string) (*Transaction, error) {
	transaction, err := usecase.Repo.GetTransaction(id)
	return transaction, err
}

func (usecase Usecase) CreateTransaction(req *Transaction) (*Transaction, error) {
	var transaction Transaction

	var total int
	// var stock int
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
		total += item.Quantity * product.Price
		// fmt.Println("total :", total)
		req.Items[i].Price = product.Price
		// req.Items[i].Product = *product
		err = usecase.ProductRepo.UpdateProduct(item.ProductID, product)
		if err != nil {
			return nil, err
		}
	}
	transaction.AdminID = req.AdminID
	transaction.Timestamp = time.Now()
	transaction.Total = total
	transaction.Items = req.Items
	fmt.Println("isi :", transaction.Items)

	err := usecase.Repo.CreateTransaction(&transaction)
	return &transaction, err
}

// 	items := []TransactionsItem{}
// 	totalPrice := 0

// 	for _, i := range req.Items {
// 		product, err := usecase.ProductRepo.GetProduct(string(i.ProductID))
// 		if err != nil {
// 			return nil, fmt.Errorf("product id not found %s", i.ProductID)
// 		}

// 		if i.Quantity > product.Stock {
// 			return nil, fmt.Errorf("stock not enough %s", product.Name)
// 		}

// 		price := int(i.Quantity) * product.Price

// 		item := &TransactionsItem{
// 			ProductID: int(i.ProductID),
// 			Quantity:  i.Quantity,
// 			Price:     price,
// 		}

// 		items = append(items, *item)

// 		totalPrice += price
// 		product.Stock = product.Stock - i.Quantity

// 		err = usecase.ProductRepo.UpdateProduct(string(i.ProductID), product)
// 		if err != nil {
// 			return nil, fmt.Errorf("data can't change")
// 		}
// 	}

// 	transaction := &Transaction{
// 		Timestamp: time.Now(),
// 		Total:     totalPrice,
// 		Items:     items,
// 	}

// 	err := usecase.Repo.CreateTransaction(transaction)
// 	if err != nil {
// 		return nil, fmt.Errorf("data can't added")
// 	}
// 		fmt.Println(string(transaction.ID))
// 	newTransaction, err := usecase.Repo.GetTransaction(string(transaction.ID))
// 	if err != nil {
// 		return nil, fmt.Errorf("error")
// 	}

// 	return newTransaction, nil
// }
