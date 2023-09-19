package models

import "time"

type Post struct {
	ID         int        `db:"id" json:"id"`
	UserID     int        `db:"user_id" json:"user_id"`
	S3Key      string     `db:"s3_key" json:"s3_key"`
	Media      MediaType  `db:"s3_media_type" json:"s3_media_type"`
	Title      string     `db:"title" json:"title"`
	Game       string     `db:"game" json:"game"`
	Body       string     `db:"body" json:"body"`
	Status     StatusType `db:"status" json:"status"`
	Likes      int        `db:"likes" json:"likes"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	LikesCount int        `json:"likes_count"`
}
