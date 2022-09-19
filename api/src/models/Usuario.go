package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa os dados do usuário utilizando a aplicação
type Usuario struct {
	ID       int       `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {

	if erro := usuario.validar(); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar() error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode ficar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode ficar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode ficar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatória e não pode ficar em branco")
	}

	return nil
}
