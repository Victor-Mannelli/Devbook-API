package models

import (
	"time"
)

type Followers struct {
	UserId     uint64    `json:"user_id,omitempty"`
	FollowerId uint64    `json:"follower_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
