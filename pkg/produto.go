package pkg

// Produto representa um item com nome e preço
type Produto struct {
	ID    int     `json:"id,omitempty"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// NovoProduto é um "construtor" idiomático
func NovoProduto(nome string, preco float64) Produto {
	return Produto{Nome: nome, Preco: preco}
}

// GetNome retorna o nome do produto
func (p Produto) GetNome() string {
	return p.Nome
}

// GetPreco retorna o preço do produto
func (p Produto) GetPreco() float64 {
	return p.Preco
}
