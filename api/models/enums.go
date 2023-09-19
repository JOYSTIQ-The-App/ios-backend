package models

// RoleType enum
type RoleType string

const (
	Admin RoleType = "admin"
	Org   RoleType = "org"
	Gamer RoleType = "gamer"
)

// StatusType enum
type StatusType string

const (
	Live     StatusType = "live"
	Draft    StatusType = "draft"
	Archived StatusType = "archived"
	Reported StatusType = "reported"
)

// MediaType enum
type MediaType string

const (
	None  MediaType = "none"
	Photo MediaType = "photo"
	Video MediaType = "video"
)

// AccountType enum
type AccountType string

const (
	Public  AccountType = "public"
	Private AccountType = "Private"
	Hidden  AccountType = "hidden"
)
