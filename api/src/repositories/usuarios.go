package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}


func (u Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	stmt, err := u.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (u Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick,
		nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (u Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := u.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
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
