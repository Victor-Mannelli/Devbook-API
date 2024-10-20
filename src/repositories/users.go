package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (usersRepository users) CreateUser(createUserDto models.User) (uint64, error) {
	statement, err := usersRepository.db.Prepare(
		"insert into users (name, email, username, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		createUserDto.Name, createUserDto.Email, createUserDto.Username, createUserDto.Password,
	)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}

func (usersRepository users) FindFilteredUsers(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // returns %nameOrUsername% which is a format needed for this query
	lines, err := usersRepository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Created_At,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (usersRepository users) FindUser(userId uint64) (models.User, error) {
	lines, err := usersRepository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE id = ?",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Created_At,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (usersRepository users) UpdateUser(userId uint64, updatedUserDto models.User) error {
	statement, err := usersRepository.db.Prepare(
		"UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(updatedUserDto.Name, updatedUserDto.Username, updatedUserDto.Email, userId); err != nil {
		return err
	}

	return nil
}
