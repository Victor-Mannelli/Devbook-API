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

func (postsRepository posts) CreatePost(createpostDto models.Posts) (uint64, error) {
	statement, err := postsRepository.db.Prepare(
		"insert into posts () values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
	//
	)

	if err != nil {
		return 0, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(postId), nil
}

func (postsRepository posts) FindFilteredposts(nameOrpostname string) ([]models.Posts, error) {
	// nameOrpostname = fmt.Sprintf("%%%s%%", nameOrpostname) // returns %nameOrpostname% which is a format needed for this query
	// rows, err := postsRepository.db.Query(
	// 	"SELECT id, name, postname, email, created_at FROM posts WHERE name LIKE ? OR postname LIKE ?",
	// 	nameOrpostname, nameOrpostname,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var posts []models.Posts
	// for rows.Next() {
	// 	var post models.Posts
	// 	if err = rows.Scan(
	// 		&post.ID,
	// 		&post.Name,
	// 		&post.postname,
	// 		&post.Email,
	// 		&post.CreatedAt,
	// 	); err != nil {
	// 		return nil, err
	// 	}
	// 	posts = append(posts, post)
	// }
	// return posts, nil
}

func (postsRepository posts) FindPostById(postId uint64) (models.Posts, error) {
	rows, err := postsRepository.db.Query(
		"SELECT id, name, postname, email, created_at FROM posts WHERE id = ?",
		postId,
	)
	if err != nil {
		return models.Posts{}, err
	}

	var post models.Posts
	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Name,
			&post.postname,
			&post.Email,
			&post.CreatedAt,
		); err != nil {
			return models.Posts{}, err
		}
	}
	return post, nil
}

func (postsRepository posts) UpdatePost(postId uint64, updatedpostDto models.Posts) error {
	query := "UPDATE posts SET "
	args := []interface{}{}

	if updatedpostDto.Name != "" {
		query += "name = ?, "
		args = append(args, updatedpostDto.Name)
	}
	if updatedpostDto.Email != "" {
		query += "email = ?, "
		args = append(args, updatedpostDto.Email)
	}
	if updatedpostDto.postname != "" {
		query += "postname = ?, "
		args = append(args, updatedpostDto.postname)
	}

	// Remove the trailing comma and space from the query
	query = strings.TrimSuffix(query, ", ")
	// Add the WHERE clause
	query += " WHERE id = ?"
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
	return nil
}
