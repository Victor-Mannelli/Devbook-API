package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	PostId         uint64    `json:"post_id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorId       uint64    `json:"author_id,omitempty"`
	AuthorUsername string    `json:"author_postname,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func (post *Post) ParsePostDto(step string) error {
	if err := post.validatePostDto(step); err != nil {
		return err
	}
	post.FormatPost()

	return nil
}

func (post *Post) validatePostDto(step string) error {
	if step == "createPost" {
		if post.Title == "" {
			return errors.New("title is required")
		}

		if post.Content == "" {
			return errors.New("content is required")
		}
	}

	if step == "updatePost" {
		if post.Title == "" && post.Content == "" {
			return errors.New("at least one field is required")
		}
	}

	return nil
}

func (post *Post) FormatPost() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
