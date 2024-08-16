package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		slog.Error("|", err)
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		slog.Error("|", err)
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		slog.Error("|", err)
		w.Write([]byte("Erro ao converter ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		slog.Error("FAILED TO CREATE USER |", err)
		w.Write([]byte("Erro ao criar o statement"))
	}
	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		slog.Error("FAILED TO CREATE USER |", err)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		slog.Error("FAILED TO GET LAST ID |", err)
		w.Write([]byte("Erro ao obter o ID inserido!"))
	}

	// STATUS CODES
	slog.Info("SUCCESS | ID", idInserido)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário cadastrado com sucesso! ID: %d", idInserido)))

}

// BuscarUsuarios -> busca usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		slog.Error("FAILED TO CONNECT DATABASE |", err)
		w.Write([]byte(fmt.Sprintf("Erro ao conectar ao banco de dados!", err)))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		slog.Error("FAILED TO CONSULT DATABASE |", err)
		w.Write([]byte("Erro ao consultar o banco de dados!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			slog.Error("FAILED TO SCAN USER", err)
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		slog.Error("FAILED TO CONVERT USER TO JSON |", err)
		w.Write([]byte("Erro ao converter os dados para Json!"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		slog.Error("FAILED CONVERT PARAM TO INT |", err)
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		slog.Error("FAILED TO DATABASE CONNECT |", err)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		slog.Error("FAILED TO SELECT USER |", err)
		w.Write([]byte("Erro ao consultar usuário no banco de dados!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			slog.Error("FAILED TO SCAN |", err)
			w.Write([]byte("Erro ao escanear os dados!"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		slog.Error("FAILED ENCODE TO JSON DATA |", err)
		w.Write([]byte("Erro ao converter para Json!"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		slog.Error("FAILED CONVERT PARAM TO INT |", err)
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	corpoReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		slog.Error("FAILED TO READ A REQUEST BODY |", err)
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(corpoReq, &usuario); err != nil {
		slog.Error("FAILED TO CONVERT A REQUEST BODY TO STRUCT |", err)
		w.Write([]byte("Erro ao converter o usuario para struct!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		slog.Error("FAILED TO CONNECT TO DATABASE |", err)
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		slog.Error("FAILED TO CREATE STATEMENT |", err)
		w.Write([]byte("Falha ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		slog.Error("FAILED TO UPDATE IN THE DATABASE |", err)
		w.Write([]byte("Falha ao atualizar o registro no banco de dados!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("ID selecionado: %d", ID)))
	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		slog.Error("FAILED CONVERT PARAM TO INT |", err)
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		slog.Error("FAILED TO CONNECT TO DATABASE |", err)
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		slog.Error("FAILED TO CREATE STATEMENT |", err)
		w.Write([]byte("Falha ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		slog.Error("FAILED TO DELETE USER")
		w.Write([]byte("Falha ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
