package banco

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golangg:golangg@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	fmt.Println("Conexão está aberta!")

	return db, nil
}
