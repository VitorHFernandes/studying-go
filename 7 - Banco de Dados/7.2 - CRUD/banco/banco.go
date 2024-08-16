package banco

import (
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		slog.Error("[RETURN]", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		slog.Error("[RETURN]", err)
		return nil, err
	}

	return db, nil
}
