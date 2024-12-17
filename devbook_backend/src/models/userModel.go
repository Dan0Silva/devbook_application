package models

import (
	"devbook_backend/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

type UpdateUserPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (u *User) Prepare(step string) error {
	if err := u.validateFields(step); err != nil {
		return err
	}

	if err := u.formateFields(step); err != nil {
		return err
	}
	return nil
}

func (u User) validateFields(step string) error {
	if u.Name == "" {
		return errors.New("mandatory field name is not filled in")
	}

	if u.Nick == "" {
		return errors.New("mandatory field nick is not filled in")
	}

	if u.Email == "" {
		return errors.New("mandatory field email is not filled in")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email format")
	}

	if step == "register" && u.Password == "" {
		return errors.New("mandatory field password is not filled in")
	}

	return nil
}

func (u *User) formateFields(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return errors.New("error performing hash operation on password")
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
