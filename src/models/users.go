package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
}

func (user *User) ParseUserDto() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.trimParser()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name is set as required")
	}
	if user.Email == "" {
		return errors.New("email is set as required")
	}
	if user.Username == "" {
		return errors.New("username is set as required")
	}
	if user.Password == "" {
		return errors.New("password is set as required")
	}

	return nil
}

func (user *User) trimParser() {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}