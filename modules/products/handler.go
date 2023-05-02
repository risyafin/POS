package products

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := handler.Usecase.GetProducts()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	_, err = json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respon := Respon{Massage: "Succes", Data: products}
	hasil, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)

}
