package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) SearchingProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	keywords := r.URL.Query().Get("keywords")
	Product, err := handler.Usecase.SearchingProduct(keywords)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = json.Marshal(Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respons := Respons{Message: "Succes", Data: Product}
	hasil, err := json.Marshal(respons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)

}

func (handler Handler) Shorting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	columStr := r.URL.Query().Get("colum")
	if columStr == "" {
		columStr = "name"
	}

	products, err := handler.Usecase.Shorting(columStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, errJson := json.Marshal(products)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusInternalServerError)
		return
	}

	respon := Respons{Message: "Succes", Data: products}
	hasil, errRespon := json.Marshal(respon)
	if errJson != nil {
		http.Error(w, errRespon.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	columStr := r.URL.Query().Get("colum")
	if columStr == "" {
		columStr = "id"
	}

	sortedStr := r.URL.Query().Get("sorted")
	if sortedStr == "" {
		sortedStr = "ASC"
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offsetStr := r.URL.Query().Get("skip")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	keywords := r.URL.Query().Get("keywords")

	products, err := handler.Usecase.GetProducts(limit, offset, columStr, sortedStr, keywords)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respons := Respons{Message: "succes", Data: products}
	hasil, err := json.Marshal(respons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("ini", r)

	product, err := handler.Usecase.GetProduct(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product *Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.Usecase.CreateProduct(product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	respon := Respons{Message: "Succes", Data: []Product{*product}}
	hasil, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := handler.Usecase.SoftDelete(id, req.Status)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	respon := Respons{Message: "Succes", Data: []Product{*data}}
	hasil, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	fmt.Println(idInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product *Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = handler.Usecase.UpdateProduct(idInt, product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	respon := Respons{Message: "Succes", Data: []Product{*product}}
	hasil, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)

}
