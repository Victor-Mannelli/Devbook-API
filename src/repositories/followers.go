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

func (followersRepository followers) FindFollowers(userId uint64) ([]models.Followers, error) {
	rows, err := followersRepository.db.Query(
		"SELECT * FROM followers WHERE user_id = ?",
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.Followers
	for rows.Next() {
		var follower models.Followers
		if err = rows.Scan(
			&follower.UserId,
			&follower.FollowerId,
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
