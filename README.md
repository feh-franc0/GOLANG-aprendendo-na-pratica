# My Go Project

This is a simple Go project that demonstrates the structure and organization of a Go application.

## Project Structure

```
my-go-project
├── cmd
│   └── main.go        # Entry point of the application
├── internal
│   └── example.go     # Internal logic and functionality
├── pkg
│   └── example.go     # Reusable code for external use
├── go.mod             # Module definition file
└── go.sum             # Checksums for module dependencies
```

## Getting Started

To get started with this project, follow these steps:

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd my-go-project
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

Provide instructions on how to use the application here.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.


## Next Steps

1️⃣ Fundamentos     → Variáveis, funções, loops
2️⃣ Structs         → Dados + métodos (como classes)
3️⃣ Organização     → Pacotes, go.mod, pastas
4️⃣ Arquivos        → Ler, salvar, logs
5️⃣ Concorrência    → Goroutines, Channels
6️⃣ HTTP/API        → Criar e consumir
7️⃣ Testes & Erros  → `error`, `testing`
8️⃣ Banco de Dados  → `sql`, persistência
9️⃣ Boas práticas   → Formatadores, linter, build



## CURL PARA API:

GET
```
   curl -X GET http://localhost:8080/usuarios
```

POST
```
   curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{"nome": "Ana Silva", "email": "ana@email.com"}'
```

PUT
```
   curl -X PUT http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "nome": "Ana Atualizada", "email": "ana@novoemail.com"}'
```

DELETE
```
   curl -X DELETE http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```
