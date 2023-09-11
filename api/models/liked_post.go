package models

type LikedPost struct {
	UserID int `db:"user_id" json:"user_id"`
	PostID int `db:"post_id" json:"post_id"`
}
