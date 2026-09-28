package models

import (
	"time"
)

type Profile struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Resume    string    `db:"resume" json:"resume"`
	Bio       string    `db:"bio" json:"bio"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
