package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, mundo!")
}

func main() {
	// Esta é apenas uma configuração mínima para ilustração.
	// Você precisará adicionar a lógica necessária para lidar com suas rotas e funcionalidades.

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
