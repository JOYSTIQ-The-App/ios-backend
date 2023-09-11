package models

// RoleType enum
type RoleType string

const (
	Admin RoleType = "admin"
	Org   RoleType = "org"
	Gamer RoleType = "gamer"
)

// Status enum
type Status string

const (
	Live     Status = "live"
	Archived Status = "archived"
	Reported Status = "reported"
)
