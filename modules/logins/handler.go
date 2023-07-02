package logins

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var admin *Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.Usecase.Registration(admin)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	json.NewDecoder(r.Body).Decode(&admin)
	token, err := handler.Usecase.Login(admin.Username, admin.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Write([]byte("Login Failed"))
		} else {
			w.Write([]byte(err.Error()))
		}
		return
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("token :", token)
	w.Write([]byte(token))
}
