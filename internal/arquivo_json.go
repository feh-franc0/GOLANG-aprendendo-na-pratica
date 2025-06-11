package internal

import (
	"encoding/json"
	"os"

	"my-go-project/pkg"
)

const caminhoJSON = "produtos.json"

// Salva os produtos no formato JSON
func SalvarProdutosJSON(produtos []pkg.Produto) error {
	bytes, err := json.MarshalIndent(produtos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(caminhoJSON, bytes, 0644)
}

// Lê produtos do JSON
func LerProdutosJSON() ([]pkg.Produto, error) {
	bytes, err := os.ReadFile(caminhoJSON)
	if err != nil {
		return nil, err
	}

	var produtos []pkg.Produto
	err = json.Unmarshal(bytes, &produtos)
	return produtos, err
}
