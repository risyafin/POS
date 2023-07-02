package main

import (
	"fmt"
	"net/http"
	"store/modules/logins"
	"store/modules/products"
	"store/modules/store"
	"store/modules/transaction"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/pos?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}

	LoginRepo := logins.Repository{DB: db}
	LoginUsecase := logins.Usecase{Repo: LoginRepo}
	LoginHandler := logins.Handler{Usecase: LoginUsecase}

	BrandsRepo := store.Repository{DB: db}
	BrandsUsecase := store.Usecase{Repo: BrandsRepo}
	BrandsHandler := store.Handler{Usecase: BrandsUsecase}

	ProductRepo := products.Repository{DB: db}
	ProductUsecase := products.Usecase{Repo: ProductRepo}
	ProductHandler := products.Handler{Usecase: ProductUsecase}

	TransactionRepo := transaction.Repository{DB: db}
	TransactionUsecase := transaction.Usecase{Repo: TransactionRepo, ProductRepo: ProductRepo}
	transactionHandler := transaction.Handler{Usecase: TransactionUsecase}

	const port string = ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/registration", LoginHandler.Registration).Methods("POST")
	r.HandleFunc("/logins", LoginHandler.Login).Methods("POST")
	r.HandleFunc("/brands", BrandsHandler.GetBrands).Methods("GET")
	r.HandleFunc("/products", jwtMiddleware(ProductHandler.GetProducts)).Methods("GET")
	r.HandleFunc("/products/searching", jwtMiddleware(ProductHandler.SearchingProduct)).Methods("GET")
	// r.HandleFunc("/products", jwtMiddleware(ProductHandler.GetProducts)).Methods("GET")
	r.HandleFunc("/products/shorting", ProductHandler.Shorting).Methods("GET")
	r.HandleFunc("/products", jwtMiddleware(ProductHandler.CreateProduct)).Methods("POST")
	r.HandleFunc("/products/{id}", jwtMiddleware(ProductHandler.GetProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(ProductHandler.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id}/status", jwtMiddleware(ProductHandler.SoftDelete)).Methods("PATCH")
	// r.HandleFunc("/products/{id}/restore", ProductHandler.RestoreProduct).Methods("PATCH")
	r.HandleFunc("/transactions", jwtMiddleware(transactionHandler.GetTransactions)).Methods("GET")
	r.HandleFunc("/transactions/{id}", jwtMiddleware(transactionHandler.GetTransaction)).Methods("GET")
	r.HandleFunc("/transactions", jwtMiddleware(transactionHandler.CreateTransaction)).Methods("POST")
	fmt.Println("localhost:8080")
	http.ListenAndServe(port, r)
}
