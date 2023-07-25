package transaction

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
	colum := r.URL.Query().Get("colum")
	if colum == "" {
		colum = "id"
	}
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "ASC"
	}
	search := r.URL.Query().Get("search")

	Transaction, err := handler.Usecase.GetTransactions(limit, offset, colum, sort, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	fmt.Println("ini ID dari context", adminId)
	if !ok {
		http.Error(w, errorAdminID.Error(), http.StatusInternalServerError)
		return
	}
	branchId, ok := r.Context().Value("branchId").(int)
	if !ok {
		http.Error(w, errorBranchID.Error(), http.StatusInternalServerError)
		return
	}
	adminUser, ok := r.Context().Value("username").(string)
	fmt.Println("ini username dari context:", adminUser)
	if !ok {
		http.Error(w, errorUsername.Error(), http.StatusInternalServerError)

		return
	}
	request.AdminID = adminId
	request.BranchID = branchId
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
