package autorisation

import (
	"encoding/json"
	"net/http"
	"todo/internal/sql"
)

type registr struct {
	email    string `json:email`
	password string `json:password`
}

func Reg(w http.ResponseWriter, r *http.Request) {
	var req registr
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.email == "" || req.password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}
	err := sql.RegDb(req.email, req.password)
	if err != nil {
		http.Error(w, "Email exist", http.StatusBadRequest)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}
