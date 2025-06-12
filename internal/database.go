package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Conectar() {
	var err error
	DB, err = sql.Open("sqlite3", "produtos.db")
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	// Cria tabela se não existir
	sqlTable := `
	CREATE TABLE IF NOT EXISTS produtos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT,
		preco REAL
	);
	`
	_, err = DB.Exec(sqlTable)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
}
