package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MYSQL
	StringConexaoBanco = ""
	// Porta onde a API vai estar rodando
	Porta = 0
)

// Carregar vai iniciar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	portDb, erro := strconv.Atoi(os.Getenv("DB_PORT"))
	if erro != nil{
		portDb = 3306
	}

	StringConexaoBanco = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		portDb,
		os.Getenv("DB_DATABASE"),
	)

}
