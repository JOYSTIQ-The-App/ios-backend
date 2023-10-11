package validators

import (
	"regexp"
)

// IsValidUsername checks if the given username meets the specified criteria
func IsValidUsername(username string) bool {
	// Regex to validate username
	const pattern = "^[a-zA-Z][a-zA-Z0-9_-]{2,14}$"

	matched, _ := regexp.MatchString(pattern, username)
	return matched
}
