package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/handlers"
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

	// Adicione o manipulador CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Ou especifique a origem desejada
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "apikey"}),
	)(r)

	http.Handle("/", corsHandler)

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
	// Aqui você pode adicionar a lógica específica para o seu caso

	// Vamos supor que você já tenha uma função para salvar o usuário no banco de dados
	saveUserToDatabase(user)

	// Agora, vamos enviar os dados para outra API
	err = sendDataToAnotherAPI(user)
	if err != nil {
		http.Error(w, "Erro ao enviar dados para outra API", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(user)
}

func saveUserToDatabase(user User) {
	// Lógica para salvar o usuário no banco de dados
	// Implemente de acordo com suas necessidades
}

func sendDataToAnotherAPI(user User) error {
	// Configuração do cliente HTTP
	client := &http.Client{}

	// URL da segunda API
	apiURL := "https://sbsmidvdvyvweqfdswee.supabase.co/rest/v1/Users"

	// Codifica o usuário em JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Cria uma solicitação POST para a segunda API
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(userJSON))
	if err != nil {
		return err
	}

	// Define o cabeçalho da solicitação
	req.Header.Set("Content-Type", "application/json")

	// Adiciona o cabeçalho apikey à solicitação
	req.Header.Set("apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InNic21pZHZkdnl2d2VxZmRzd2VlIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDEzMDA4NzYsImV4cCI6MjAxNjg3Njg3Nn0.559lUmQSZd8sexAP4xtBBPkm8EN93oMVn6wy4jCFPIs")

	// Executa a solicitação
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verifica o status da resposta da segunda API
	if resp.StatusCode != http.StatusCreated {
		return errors.New("Resposta inesperada da segunda API")
	}

	return nil
}
