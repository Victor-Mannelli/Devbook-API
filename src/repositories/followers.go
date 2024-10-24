package repositories

import (
	"api/src/models"
	"database/sql"
)

type followers struct {
	db *sql.DB
}

func FollowersRepository(db *sql.DB) *followers {
	return &followers{db}
}

func (followersRepository users) FindFollowers(userId uint64) ([]models.User, error) {
	rows, err := followersRepository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created_at
		FROM users u INNER JOIN followers f on u.id = f.follower_id
		WHERE user_id = ?
	`, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User
		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Username,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func (followersRepository users) FindFollowing(userId uint64) ([]models.User, error) {
	rows, err := followersRepository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created_at
		FROM users u INNER JOIN followers f on u.id = f.user_id
		WHERE follower_id = ?
	`, userId,
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (followersRepository followers) Follow(userId uint64, userToFollow uint64) error {
	statement, err := followersRepository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userToFollow, userId); err != nil {
		return err
	}

	return nil
}

func (followersRepository followers) UnFollow(followerId uint64, userId uint64) error {
	statement, err := followersRepository.db.Prepare("DELETE FROM followers WHERE user_id = ? && follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}
