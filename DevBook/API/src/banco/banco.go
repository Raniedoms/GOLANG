package banco

import (
	"API/API/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver
)

//Conectar abre a conexao com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
