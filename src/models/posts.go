package models

import "time"

type Posts struct {
	PostId         uint64    `json:"post_id,omitempty"`
	Title          uint64    `json:"title,omitempty"`
	Content        uint64    `json:"content,omitempty"`
	AuthorId       uint64    `json:"author_id,omitempty"`
	AuthorUsername uint64    `json:"author_username,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
