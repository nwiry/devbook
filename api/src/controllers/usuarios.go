package controllers

import (
	"api/responses"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	enteredUser, err := repositorio.Criar(usuario)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	usuario.ID = int(enteredUser)

	responses.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario busca um usuário salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}

// Altera as informações do usuario salvo no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

}

// DeletarUsuario deleta um usuário salvo no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

}
