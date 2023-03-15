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
	line, err := p.db.Query(
		"select p.*, u.nick from publicacoes p "+
			"inner join usuarios u on u.id = p.autor_id "+
			"where p.id = ?", ID)

	if err != nil {
		return domain.Publish{}, err
	}

	defer line.Close()

	var publish domain.Publish

	for line.Next() {
		if err := line.Scan(
			&publish.ID,
			&publish.Title,
			&publish.Content,
			&publish.AuthorID,
			&publish.Likes,
			&publish.DateCreate,
			&publish.AuthorNick); err != nil {
			return publish, err
		}
	}
	return publish, nil
}

func (p publish) GetAllPublish(ID uint64) ([]domain.Publish, error) {
	lines, err := p.db.Query("select distinct p.*, u.nick from publicacoes p "+
		"inner join usuarios u on u.id = p.autor_id "+
		"inner join seguidores s on p.autor_id = s.usuario_id "+
		"where u.id = ? or s.seguidor_id = ?", ID, ID)

	if err != nil {
		return []domain.Publish{}, err
	}

	defer lines.Close()

	var publishs []domain.Publish
	for lines.Next() {
		var publish domain.Publish
		if err := lines.Scan(
			&publish.ID,
			&publish.Title,
			&publish.Content,
			&publish.AuthorID,
			&publish.Likes,
			&publish.DateCreate,
			&publish.AuthorNick); err != nil {
			return publishs, err
		}
		publishs = append(publishs, publish)
	}

	return publishs, nil
}
