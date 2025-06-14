package main

import (
	"log"
	"my-go-project/internal"
	"net/http"
)

func main() {
	internal.Conectar()

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			internal.ListarUsuariosHandler(w, r)
		} else if r.Method == http.MethodPost {
			internal.CriarUsuarioHandler(w, r)
		} else if r.Method == http.MethodPut {
			internal.AtualizarUsuarioHandler(w, r)
		} else if r.Method == http.MethodDelete {
			internal.DeletarUsuarioHandler(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
