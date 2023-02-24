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

func (usersRepository users) GetAll() ([]domain.Users, error) {
	lines, err := usersRepository.db.Query("select nome, nick, email from usuarios")

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []domain.Users
	for lines.Next() {
		var user domain.Users
		if err := lines.Scan(&user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (usersRepository users) GetUserId(ID uint64) (domain.Users, error) {
	line, err := usersRepository.db.Query("select nome, nick, email from usuarios where id=?", ID)

	if err != nil {
		return domain.Users{}, err
	}

	var user domain.Users

	for line.Next() {
		if err := line.Scan(&user.Name, &user.Nick, &user.Email); err != nil {
			return domain.Users{}, err
		}
	}

	return user, nil
}
