package models

import (
	"errors"
	"strings"
	"time"
)

// * User representa um usuário utilizando a rede social.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pass      string    `json:"pass,omitempty"`
	CreatedAt time.Time `json:"dtCreated,omitempty"`
}

// * Prepare chama os métodos para validar e formatar os usuários recebidos.
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome e obrigatorio e nao pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nick e obrigatorio e nao pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail e obrigatorio e nao pode estar em branco")
	}

	if step == "register" && user.Pass == "" {
		return errors.New("A senha e obrigatoria e nao pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
