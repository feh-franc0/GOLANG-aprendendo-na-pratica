package internal

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Conectar() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/golangdb")
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	criarTabelaUsuarios()
}

func criarTabelaUsuarios() {
	query := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nome VARCHAR(100),
		email VARCHAR(100)
	);`

	if _, err := DB.Exec(query); err != nil {
		log.Fatal("Erro ao criar tabela usuarios:", err)
	}
}
