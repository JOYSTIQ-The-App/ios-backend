package models

import "time"

type User struct {
	ID          int         `db:"id" json:"id"`
	AccountType AccountType `db:"account_type" json:"account_type"`
	Email       string      `db:"email" json:"email"`
	Username    string      `db:"username" json:"username"`
	Role        RoleType    `db:"role" json:"role"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
}
