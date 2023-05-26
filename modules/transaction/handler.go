package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	transaction, err := handler.Usecase.GetTransaction(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(hasil)
}

func (handler Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	Transaction, err := handler.Usecase.GetTransactions()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = json.Marshal(Transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respon := GetAllTransactionsResponse{Message: "Succes", Data: Transaction}
	hasil, err := json.Marshal(respon)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(hasil)
}

func (handler Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("idnya adalah ", r.Context().Value("adminId"))

	var request *Transaction

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	adminId, ok := r.Context().Value("adminId").(int)
	if !ok {
		errors.New("adminID not int ")
		return
	}
	adminUser, ok := r.Context().Value("username").(string)
	fmt.Println("ini username :", adminUser)
	if !ok {
		errors.New("username not string")
		return
	}
	request.AdminID = adminId
	request.Admin.Username = adminUser
	fmt.Println(request.Admin.Username)
	transaction, err := handler.Usecase.CreateTransaction(request)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	respon := GetDeteailTransactionResponse{Message: "Succes", Data: transaction}
	hasil, err := json.Marshal(respon)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(hasil)
}
