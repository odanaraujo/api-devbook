package domain

import (
	"errors"
	"strings"
	"time"
)

type Publish struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick uint64    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	DateCreate time.Time `json:"date_create,omitempty"`
}

func (publish *Publish) Prepare() error {
	if err := publish.validator(); err != nil {
		return err
	}
	publish.formater()
	return nil
}

func (publish *Publish) validator() error {
	if publish.Title == "" {
		return errors.New("The title is required")
	}
	if publish.Content == "" {
		return errors.New("The content is required")
	}

	return nil
}

func (publish *Publish) formater() {
	publish.Title = strings.TrimSpace(publish.Title)
	publish.Content = strings.TrimSpace(publish.Content)
}
