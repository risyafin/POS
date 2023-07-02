package branch

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetBrands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	skipStr := r.URL.Query().Get("skip")
	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		skip = 0
	}

	columStr := r.URL.Query().Get("colum")
	if columStr == "" {
		columStr = "id"
	}

	sortStr := r.URL.Query().Get("sort")
	if sortStr == "" {
		sortStr = "ASC"
	}

	search := r.URL.Query().Get("search")

	brands, err := handler.Usecase.GetBrands(limit, skip, columStr, sortStr, search)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = json.Marshal(brands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respon := Respon{Message: "Succes", Data: brands}
	hasil, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}
