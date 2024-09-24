package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// * Users representa um repositório de usuários.
type Users struct {
	db *sql.DB
}

// * NewUserRepository recebe um DB vindo do controller de usuários e aplica no struct de usuário.
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// * Create insere um usuário no banco de dados.
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, pass) values(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Pass)
	if err != nil {
		return 0, nil
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

// * Find trás todos os usuários que atendem ao filtro de nome ou nick, Atribuídos a variável n.
func (repository Users) Find(n string) ([]models.User, error) {
	n = fmt.Sprintf("%%%s%%", n) //* %nomeOuNick%

	line, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users WHERE name LIKE ? or nick LIKE ?",
		n, n,
	)
	if err != nil {
		return nil, err
	}
	defer line.Close()

	var users []models.User

	for line.Next() {
		var user models.User
		if err = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindById(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}
