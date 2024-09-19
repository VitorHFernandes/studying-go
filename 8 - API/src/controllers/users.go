package controllers

import "net/http"

//* Insere um usuário no banco de dados.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

//* Busca todos os usuários salvos no banco de dados.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar todos os usuário!"))
}

//* Busca um único usuário por ID.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um único usuário por ID!"))
}

//* Altera as informações de um usuário no banco.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um único usuário por ID!"))
}

//* Excluí as informações de um usuário no banco.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta usuario"))
}
