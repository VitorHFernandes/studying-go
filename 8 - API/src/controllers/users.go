package controllers

import (
	"log/slog"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("| CREATING USER")
	w.Write([]byte("Criando usuário"))
	w.WriteHeader(http.StatusCreated)
}

func SelectUsers(w http.ResponseWriter, r *http.Request) {
	slog.Info("| SELECTING USERS")
	w.Write([]byte("Buscando todos os usuários"))
	w.WriteHeader(http.StatusOK)
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("| SELECTING USER")
	w.Write([]byte("Buscando um único usuário"))
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("| UPDATING USER")
	w.Write([]byte("Atualizando um usuário"))
	w.WriteHeader(http.StatusAccepted)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("| DELETING USER")
	w.Write([]byte("Deletando um usuário"))
	w.WriteHeader(http.StatusNoContent)
}
