package main

import (
	"fmt"
	"net/http"
	"store/modules/products"
	"store/modules/transaction"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/pos"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	ProductRepo := products.Repository{DB: db}
	ProductUsecase := products.Usecase{Repo: ProductRepo}
	ProductHandler := products.Handler{Usecase: ProductUsecase}

	TransactionRepo := transaction.Repository{DB: db}
	TransactionUsecase := transaction.Usecase{Repo: TransactionRepo}
	transactionHandler := transaction.Handler{Usecase: TransactionUsecase}

	const port string = ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/products", ProductHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products", ProductHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", ProductHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", ProductHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/transactions", transactionHandler.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", transactionHandler.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")
	fmt.Println("lohalhost:8080")
	http.ListenAndServe(port, r)
}
