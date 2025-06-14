package internal

import (
	"encoding/json"
	"my-go-project/pkg"
	"net/http"
)

func CriarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var u pkg.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	if err := criarUsuario(u); err != nil {
		http.Error(w, "Erro ao salvar", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ListarUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	usuarios, err := listarUsuarios()
	if err != nil {
		http.Error(w, "Erro ao listar", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usuarios)
}

func AtualizarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var u pkg.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	if err := atualizarUsuario(u); err != nil {
		http.Error(w, "Erro ao atualizar", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var u pkg.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	if err := deletarUsuario(u.ID); err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
