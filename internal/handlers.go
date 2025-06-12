package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"my-go-project/pkg"
)

// GET /produtos
func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, nome, preco FROM produtos")
	if err != nil {
		http.Error(w, "Erro ao consultar", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []pkg.Produto
	for rows.Next() {
		var p pkg.Produto
		err := rows.Scan(&p.ID, &p.Nome, &p.Preco)
		if err != nil {
			log.Println("Erro ao fazer scan:", err)
			continue
		}
		produtos = append(produtos, p)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(produtos); err != nil {
		http.Error(w, "Erro ao converter JSON", http.StatusInternalServerError)
	}
}

// POST /produtos
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	var p pkg.Produto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec("INSERT INTO produtos (nome, preco) VALUES (?, ?)", p.Nome, p.Preco)
	if err != nil {
		http.Error(w, "Erro ao salvar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto %s criado com sucesso!", p.Nome)
}
