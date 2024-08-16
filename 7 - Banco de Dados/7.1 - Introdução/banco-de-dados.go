package main

import (
	"database/sql" // ESSE PACOTE USA O ULTIMO PACOTE INSERIDO
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql" // _ USAR PACOTE DE FORMA IMPLÍCITA
)

func main() {
	// stringConexao := "usuario:senha@/banco?charset=utf8&parseTime=True&loc=Local"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		slog.Error("[RETURN]", err)
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		slog.Error("[RETURN]", err)
		panic(err)
	}
	slog.Info("| CONEXÃO ESTABELECIDA")

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		slog.Error("[RETURN]", err)
		panic(err)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
