package repositories

import (
	"api/src/models"
	"database/sql"
	"strings"
)

type posts struct {
	db *sql.DB
}

func PostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (postsRepository posts) CreatePost(createpostDto models.Post) (uint64, error) {
	statement, err := postsRepository.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(createpostDto.Title, createpostDto.Content, createpostDto.AuthorId)

	if err != nil {
		return 0, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(postId), nil
}

func (postsRepository posts) FindPosts(userId uint64) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		SELECT DISTINCT p.*, u.username 
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = ? or f.follower_id = ?
		ORDER BY 1 desc`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.PostId,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (postsRepository posts) FindPostById(postId uint64) (models.Post, error) {
	rows, err := postsRepository.db.Query(`
		SELECT p.*, u.username 
		FROM posts p INNER JOIN users u ON u.id = p.author_id 
		WHERE p.post_id = ?`,
		postId,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		if err = rows.Scan(
			&post.PostId,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

func (postsRepository posts) FindPostsByUser(userId uint64) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		SELECT p.*, u.username 
		FROM posts p INNER JOIN users u ON u.id = p.author_id 
		WHERE p.author_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.PostId,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (postsRepository posts) UpdatePost(postId uint64, updatedpostDto models.Post) error {
	query := "UPDATE posts SET "
	args := []interface{}{}

	if updatedpostDto.Title != "" {
		query += "title = ?, "
		args = append(args, updatedpostDto.Title)
	}
	if updatedpostDto.Content != "" {
		query += "content = ?, "
		args = append(args, updatedpostDto.Content)
	}

	query = strings.TrimSuffix(query, ", ") + " WHERE post_id = ?"
	args = append(args, postId)

	statement, err := postsRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(args...); err != nil {
		return err
	}

	return nil
}

func (postsRepository posts) DeletePost(postId uint64) error {
	statement, err := postsRepository.db.Prepare("DELETE FROM posts WHERE post_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
