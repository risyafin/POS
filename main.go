package main

import (
	"fmt"
	"net/http"
	"store/modules/branches"
	"store/modules/config"
	"store/modules/logins"
	"store/modules/products"
	"store/modules/transaction"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	conf := config.GetConfig()
	conn := config.InitDatabaseConnection(conf)
	config.AutoMigration(conn)

	LoginRepo := logins.Repository{DB: conn}
	LoginUsecase := logins.Usecase{Repo: LoginRepo}
	LoginHandler := logins.Handler{Usecase: LoginUsecase}

	BranchRepo := branches.Repository{DB: conn}
	BranchUsecase := branches.Usecase{Repo: BranchRepo}
	BranchHandler := branches.Handler{Usecase: BranchUsecase}

	ProductRepo := products.Repository{DB: conn}
	ProductUsecase := products.Usecase{Repo: ProductRepo}
	ProductHandler := products.Handler{Usecase: ProductUsecase}

	TransactionRepo := transaction.Repository{DB: conn}
	TransactionUsecase := transaction.Usecase{Repo: TransactionRepo, ProductRepo: ProductRepo}
	transactionHandler := transaction.Handler{Usecase: TransactionUsecase}

	const port string = ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/registration", LoginHandler.Registration).Methods("POST")
	r.HandleFunc("/logins", LoginHandler.Login).Methods("POST")
	r.HandleFunc("/branch", BranchHandler.GetBrands).Methods("GET")
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
