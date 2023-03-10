package repository

import (
	"database/sql"
	"github.com/odanaraujo/api-devbook/domain"
)

type publish struct {
	db *sql.DB
}

func NewRepositoryPublish(db *sql.DB) *publish {
	return &publish{db}
}

func (p publish) CreaterPublish(publish domain.Publish) (uint64, error) {
	statement, err := p.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	defer statement.Close()

	if err != nil {
		return 0, err
	}

	insert, err := statement.Exec(publish.Title, publish.Content, publish.AuthorID)

	if err != nil {
		return 0, err
	}

	insertID, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(insertID), nil
}

func (p publish) GetPublish(ID uint64) (domain.Publish, error) {
	line, err := p.db.Query("select titulo, conteudo, autor_id, autor_nick where id = ?", ID)

	if err != nil {
		return domain.Publish{}, err
	}

	defer line.Close()

	var publish domain.Publish

	for line.Next() {
		if err := line.Scan(&publish.Title, &publish.Content, &publish.AuthorID, publish.AuthorNick); err != nil {
			return publish, err
		}
	}
	return publish, nil
}
