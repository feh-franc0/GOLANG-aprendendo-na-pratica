package main

import (
	"fmt"
	"log"
	"net/http"

	"my-go-project/internal"
)

func main() {
	// Conecta com o banco de dados
	internal.Conectar()

	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
		// Diferencia o tipo de metodo HTTP recebido
		if r.Method == http.MethodGet {
			internal.ListarProdutos(w, r)
		} else if r.Method == http.MethodPost {
			internal.CriarProduto(w, r)
		} else {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
