package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
}

// BuscarUsuarios busca todos os usuarios salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	
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
