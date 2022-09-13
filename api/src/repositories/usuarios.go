package repositories

import (
	"api/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	stmt, err := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	lastInsert, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return uint64(lastInsert), nil
}