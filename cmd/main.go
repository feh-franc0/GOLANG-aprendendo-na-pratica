package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"my-go-project/internal"
	"my-go-project/pkg"
)

func main() {
	fmt.Println("Iniciando aplicação com goroutines e JSON")
	produtos := []pkg.Produto{
		pkg.NovoProduto("Camiseta", 49.90),
		pkg.NovoProduto("Tênis", 199.90),
		pkg.NovoProduto("Calça", 99.90),
	}

	// 🔄 Salvar no JSON
	if err := internal.SalvarProdutosJSON(produtos); err != nil {
		log.Fatal("Erro ao salvar JSON:", err)
	}
	fmt.Println("Produtos salvos em JSON.")

	// ✅ Lendo e processando com goroutines
	lidos, err := internal.LerProdutosJSON()
	if err != nil {
		log.Fatal("Erro ao ler JSON:", err)
	}

	var wg sync.WaitGroup
	resultado := make(chan string)

	// 🧵 Processamento concorrente com goroutines
	for _, p := range lidos {
		wg.Add(1)
		go func(prod pkg.Produto) {
			defer wg.Done()
			time.Sleep(1 * time.Second)          // Simula processamento
			precoFinal := prod.GetPreco() * 1.15 // 15% de imposto
			msg := fmt.Sprintf("Produto %s com imposto: R$%.2f", prod.GetNome(), precoFinal)
			resultado <- msg
		}(p)
	}

	// 🧵 Outra goroutine que coleta os resultados
	go func() {
		wg.Wait()
		close(resultado)
	}()

	// 🔄 Leitura concorrente dos resultados
	for r := range resultado {
		fmt.Println("[PROCESSADO]", r)
	}
}
