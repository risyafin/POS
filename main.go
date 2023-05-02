package main

import (
	"fmt"
	"net/http"
	"store/modules/products"

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

	const port string = ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/products", ProductHandler.GetProducts).Methods("GET")

	fmt.Println("lohalhost:8080")
	http.ListenAndServe(port, r)
}
