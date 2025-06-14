package internal

import (
	"my-go-project/pkg"
)

func criarUsuario(u pkg.Usuario) error {
	_, err := DB.Exec("INSERT INTO usuarios (nome, email) VALUES (?, ?)", u.Nome, u.Email)
	return err
}

func listarUsuarios() ([]pkg.Usuario, error) {
	rows, err := DB.Query("SELECT id, nome, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []pkg.Usuario
	for rows.Next() {
		var u pkg.Usuario
		rows.Scan(&u.ID, &u.Nome, &u.Email)
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func atualizarUsuario(u pkg.Usuario) error {
	_, err := DB.Exec("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?", u.Nome, u.Email, u.ID)
	return err
}

func deletarUsuario(id int) error {
	_, err := DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
	return err
}
