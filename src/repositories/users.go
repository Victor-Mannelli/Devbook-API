package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"strings"
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
	rows, err := usersRepository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (usersRepository users) FindUserById(userId uint64) (models.User, error) {
	rows, err := usersRepository.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE id = ?",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (usersRepository users) FindUserByEmail(email string) (models.User, error) {
	row, err := usersRepository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (usersRepository users) UpdateUser(userId uint64, updatedUserDto models.User) error {
	query := "UPDATE users SET "
	args := []interface{}{}

	if updatedUserDto.Name != "" {
		query += "name = ?, "
		args = append(args, updatedUserDto.Name)
	}
	if updatedUserDto.Email != "" {
		query += "email = ?, "
		args = append(args, updatedUserDto.Email)
	}
	if updatedUserDto.Username != "" {
		query += "username = ?, "
		args = append(args, updatedUserDto.Username)
	}

	// Remove the trailing comma and space from the query
	query = strings.TrimSuffix(query, ", ")
	// Add the WHERE clause
	query += " WHERE id = ?"
	args = append(args, userId)

	statement, err := usersRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(args...); err != nil {
		return err
	}

	return nil
}

func (usersRepository users) DeleteUser(userId uint64) error {
	statement, err := usersRepository.db.Prepare("DELETE FROM users WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId); err != nil {
		return err
	}

	return nil
}
