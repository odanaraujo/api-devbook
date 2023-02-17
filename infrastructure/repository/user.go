package repository

import (
	"database/sql"
	"github.com/odanaraujo/api-devbook/domain"
)

type users struct {
	db *sql.DB
}

// NewRepositoryUser função vai receber um banco, que será aberto no controller
// controller chama esse repositório de usuário
// Função recebe o banco e joga dentro do struct de usuario
// assim, controller vai se preocupar apenas em abrir a conexão com o banco
func NewRepositoryUser(db *sql.DB) *users {
	return &users{db}
}

func (usersRepository users) Save(user domain.Users) (uint64, error) {
	statement, err := usersRepository.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	defer statement.Close()

	if err != nil {
		return 0, err
	}

	insert, err := statement.Exec(user.Name, user.Email, user.Nick, user.Password)

	if err != nil {
		return 0, err
	}

	idInsert, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(idInsert), nil
}
