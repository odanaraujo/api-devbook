package repository

import (
	"database/sql"
	"fmt"
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

func (usersRepository users) Save(user domain.User) (uint64, error) {
	statement, err := usersRepository.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	defer statement.Close()

	if err != nil {
		return 0, err
	}

	insert, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	idInsert, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(idInsert), nil
}

func (usersRepository users) GetAll(nickOrName string) ([]domain.User, error) {
	nickOrName = fmt.Sprintf("%%%s%%", nickOrName)
	lines, err := usersRepository.db.Query("select nome, nick, email, dataCriacao from usuarios where nome LIKE ? or nick LIKE ?", nickOrName, nickOrName)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []domain.User
	for lines.Next() {
		var user domain.User
		if err := lines.Scan(&user.Name, &user.Nick, &user.Email, &user.CreateDate); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (usersRepository users) GetUserId(ID uint64) (domain.User, error) {
	line, err := usersRepository.db.Query("select id, nome, nick, email, dataCriacao from usuarios where id=?", ID)

	if err != nil {
		return domain.User{}, err
	}

	var user domain.User

	for line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreateDate); err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (usersRepository users) GetPassword(ID uint64) (string, error) {
	line, err := usersRepository.db.Query("select senha from usuarios where id=?", ID)

	if err != nil {
		return "", err
	}

	var user domain.User

	for line.Next() {
		if err := line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (usersRepository users) UpdateUser(ID uint64, user domain.User) (domain.User, error) {
	statement, err := usersRepository.db.Prepare("update usuarios set nome = ?, email = ?, nick = ? where id = ?")

	if err != nil {
		return domain.User{}, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Nick, ID)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (usersRepository users) DeleteUser(ID uint64) error {
	statement, err := usersRepository.db.Prepare("delete from usuarios where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(ID)

	if err != nil {
		return err
	}

	return nil
}

func (userRepository users) FollowUser(userID uint64, followID uint64) error {
	statement, err := userRepository.db.Prepare("insert into seguidores (usuario_id, seguidor_id) values(?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followID)

	if err != nil {
		return err
	}

	return nil
}

func (userRepository users) UnfollowUser(userID uint64, followID uint64) error {
	statement, err := userRepository.db.Prepare("delete from seguidores where usuario_id = ? AND seguidor_id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followID)

	if err != nil {
		return err
	}

	return nil
}

func (userRepository users) GetFollowersUser(userID uint64) ([]domain.User, error) {
	lines, err := userRepository.db.Query(
		"select id, nome, nick, email, dataCriacao "+
			"from usuarios u "+
			"inner join seguidores s on u.id = s.seguidor_id "+
			"where s.usuario_id = ?", userID)

	if err != nil {
		return nil, err
	}

	var users []domain.User
	for lines.Next() {
		var user domain.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepository users) GetFollowingUser(userID uint64) ([]domain.User, error) {
	lines, err := userRepository.db.Query(
		"select id, nome, nick, email, dataCriacao "+
			"from usuarios u "+
			"inner join seguidores s on u.id = s.usuario_id "+
			"where s.seguidor_id = ?", userID)

	if err != nil {
		return nil, err
	}

	var users []domain.User
	for lines.Next() {
		var user domain.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepository users) UpdatePassword(userID uint64, passwordHash string) error {
	statement, err := userRepository.db.Prepare("update usuarios set senha = ? where id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(passwordHash, userID)

	if err != nil {
		return err
	}
	return nil
}
