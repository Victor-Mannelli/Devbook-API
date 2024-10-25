package models

import (
	"api/src/utils"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) ParseUserDto(step string) error {
	if err := user.validateUserDto(step); err != nil {
		return err
	}
	if err := user.formatUserDto(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validateUserDto(step string) error {
	emailErr := checkmail.ValidateFormat(user.Email)
	if user.Email != "" && emailErr != nil {
		return errors.New("email with invalid format")
	}

	if step == "updateUser" {
		if user.Name == "" && user.Email == "" && user.Username == "" && user.Password == "" {
			return errors.New("at least one field is required")
		}
		// No need to check for non-provided fields (i.e., empty fields will not be updated)
	}

	if step == "createUser" {
		if user.Name == "" {
			return errors.New("name is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if user.Username == "" {
			return errors.New("username is required")
		}
		if user.Password == "" {
			return errors.New("password is required")
		}
	}

	return nil
}

func (user *User) formatUserDto(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if step == "createUser" {
		hashedPassword, err := utils.HashString(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
