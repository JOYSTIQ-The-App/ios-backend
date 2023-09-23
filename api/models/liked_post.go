package models

import "time"

type LikedPost struct {
	UserID    int       `db:"user_id" json:"user_id"`
	PostID    int       `db:"post_id" json:"post_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
