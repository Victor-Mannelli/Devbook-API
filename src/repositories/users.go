package repositories

import (
	"api/src/modules"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (usersRepository users) CreateUser(createUserDto modules.User) (uint64, error) {
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
