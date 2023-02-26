package domain

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/odanaraujo/api-devbook/infrastructure/security"
	"strings"
	"time"
)

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (user *User) Prepare(isUpdate bool) error {
	if err := user.validator(isUpdate); err != nil {
		return err
	}

	if err := user.format(isUpdate); err != nil {
		return err
	}
	return nil
}

func (user *User) validator(isUpdate bool) error {
	if user.Name == "" {
		return errors.New("the Name is required and cannot be blank")
	}

	if user.Nick == "" {
		return errors.New("the Nick is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("the Email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Not a valid email")
	}

	if !isUpdate && user.Password == "" {
		return errors.New("the Password is required and cannot be blank")
	}

	return nil
}

func (user *User) format(isUpdate bool) error {

	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if err := formatPassword(isUpdate, user); err != nil {
		return err
	}

	return nil
}

func formatPassword(isUpdate bool, user *User) error {
	if !isUpdate {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
