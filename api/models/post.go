package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID             int            `db:"id" json:"id"`
	UserID         int            `db:"user_id" json:"user_id"`
	S3Key          sql.NullString `db:"s3_key" json:"s3_key"`
	ThumbnailS3Key sql.NullString `db:"thumbnail_s3_key" json:"thumbnail_s3_key"`
	Media          MediaType      `db:"media" json:"media"`
	Game           string         `db:"game" json:"game"`
	Body           string         `db:"body" json:"body"`
	Status         StatusType     `db:"status" json:"status"`
	Likes          int            `db:"likes" json:"likes"`
	Comments       int            `db:"comments" json:"comments"`
}

type FeedPost struct {
	Post
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UserLiked   bool           `json:"user_liked"`
	Username    string         `json:"username"`
	AvatarS3Key sql.NullString `json:"avatar_s3_key"`
}

type LeaderboardPost struct {
	Post
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	LikesCount  int            `json:"likes_count"`
	UserLiked   bool           `json:"user_liked"`
	Username    string         `json:"username"`
	AvatarS3Key sql.NullString `json:"avatar_s3_key"`
}
