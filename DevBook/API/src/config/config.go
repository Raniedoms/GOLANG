package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco é a string de conexao com o MySQL
	StringConexaoBanco = ""
	//Porta onde a API vai estar rodando
	Porta = 0
)

//godotenv - ele le os pacotes .env
//Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	//converte de string para inteiro
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}
	//Sprintf - formatar a string pra mim
	//1° %s - usuario
	//2° %s - senha
	//3° %s - nome do banco
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	//usuario | senha | nome do banco
}
