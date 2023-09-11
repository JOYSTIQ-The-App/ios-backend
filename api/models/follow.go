package models

import "time"

type Follow struct {
	FollowingUserID int       `db:"following_user_id" json:"following_user_id"`
	FollowedUserID  int       `db:"followed_user_id" json:"followed_user_id"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
