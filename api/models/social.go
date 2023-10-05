package models

import (
	"time"
)

type Social struct {
	UserID         int        `db:"user_id" json:"user_id"`
	Social         SocialType `db:"social" json:"social"`
	SocialUsername string     `db:"social_username" json:"social_username"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
}
