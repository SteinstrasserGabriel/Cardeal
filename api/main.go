package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Admin bool   `json:"admin"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", createUserHandler).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON da solicitação para a estrutura User
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lógica de processamento (por exemplo, salvar o usuário no banco de dados)

	// https://sbsmidvdvyvweqfdswee.supabase.co/rest/v1/Users 

	// Retorna uma resposta de sucesso
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
