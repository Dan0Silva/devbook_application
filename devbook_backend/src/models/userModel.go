package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validateFields(); err != nil {
		return err
	}

	u.formateFields()
	return nil
}

func (u User) validateFields() error {
	if u.Name == "" {
		return errors.New("mandatory field name is not filled in")
	}

	if u.Nick == "" {
		return errors.New("mandatory field nick is not filled in")
	}

	if u.Email == "" {
		return errors.New("mandatory field email is not filled in")
	}

	if u.Password == "" {
		return errors.New("mandatory field password is not filled in")
	}

	return nil
}

func (u *User) formateFields() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
