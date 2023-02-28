package repository

import (
	"database/sql"
	"github.com/odanaraujo/api-devbook/request"
)

type login struct {
	db *sql.DB
}

// NewRepositoryUser função vai receber um banco, que será aberto no controller
// controller chama esse repositório de usuário
// Função recebe o banco e joga dentro do struct de usuario
// assim, controller vai se preocupar apenas em abrir a conexão com o banco
func NewRepositoryLogin(db *sql.DB) *login {
	return &login{db}
}

func (loginRequest login) GetUserWithEmail(email string) (request.LoginRequest, error) {

	line, err := loginRequest.db.Query("select email, senha from usuarios where email = ?", email)

	if err != nil {
		return request.LoginRequest{}, err
	}

	defer line.Close()

	var request request.LoginRequest
	for line.Next() {
		if err := line.Scan(&request.Email, &request.Password); err != nil {
			return request, err
		}
	}
	return request, nil
}
