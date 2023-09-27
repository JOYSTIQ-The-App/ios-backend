package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int            `db:"id" json:"id"`
	UserID    int            `db:"user_id" json:"user_id"`
	S3Key     sql.NullString `db:"s3_key" json:"s3_key"`
	Media     MediaType      `db:"media" json:"media"`
	Game      string         `db:"game" json:"game"`
	Body      string         `db:"body" json:"body"`
	Status    StatusType     `db:"status" json:"status"`
	Likes     int            `db:"likes" json:"likes"`
	Comments  int            `db:"comments" json:"comments"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
}

type UserFeedPost struct {
	Post
	UserLiked bool `json:"user_liked"`
}

type LeaderboardPost struct {
	Post
	LikesCount int  `json:"likes_count"`
	UserLiked  bool `json:"user_liked"`
}
