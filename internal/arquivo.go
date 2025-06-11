package internal

import (
	"fmt"
	"my-go-project/pkg"
	"os"
	"strings"
)

// caminho do arquivo
const caminho = "produtos.txt"

// salva um slice de produtos no arquivo
// Começa com letra maiúscula se for exportada (usada fora do pacote atual)
func SalvarProdutos(produtos []pkg.Produto) error {
	var linhas []string
	for _, p := range produtos {
		linha := fmt.Sprintf("%s;%.2f", p.GetNome(), p.GetPreco())
		linhas = append(linhas, linha)
	}

	// junta tudo e grava no arquivo
	conteudo := strings.Join(linhas, "\n")
	return os.WriteFile(caminho, []byte(conteudo), 0644) // 0644: permissão do arquivo (leitura e escrita para dono, leitura para outros)
}

// Lê os produtos do arquivo e retorna slice
func LerProdutos() ([]pkg.Produto, error) {
	bytes, err := os.ReadFile(caminho)
	if err != nil {
		return nil, err
	}

	linhas := strings.Split(string(bytes), "\n")
	var produtos []pkg.Produto

	for _, linha := range linhas {
		if linha == "" {
			continue // ignora linhas vazias
		}
		partes := strings.Split(linha, ";")
		if len(partes) != 2 {
			continue
		}
		nome := partes[0]
		var preco float64
		fmt.Sscanf(partes[1], "%f", &preco) // converte string para float64
		produto := pkg.NovoProduto(nome, preco)
		produtos = append(produtos, produto)
	}

	return produtos, nil
}
