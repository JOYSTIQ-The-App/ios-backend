package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// First, retrieve the user's ID based on the provided username
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the count of followers
		var followersCount int
		err = db.QueryRow("SELECT COUNT(*) FROM follows WHERE followed_user_id = $1", userID).Scan(&followersCount)
		if err != nil {
			log.Println("Error counting followers", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the count of users the user is following
		var followingCount int
		err = db.QueryRow("SELECT COUNT(*) FROM follows WHERE following_user_id = $1", userID).Scan(&followingCount)
		if err != nil {
			log.Println("Error counting followed users", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Now, retrieve the resume and bio for the obtained user ID
		query := sb.Select("profiles.resume", "profiles.bio").From("profiles").
			Where(sq.Eq{"user_id": userID}).
			Limit(1) // Since user ID is unique, we only expect one match

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		row := db.QueryRow(sqlQuery, args...)
		var resume, bio string
		err = row.Scan(&resume, &bio)
		if err == sql.ErrNoRows {
			resume, bio = "", "" // Set blank strings if no profile is found
		} else if err != nil {
			log.Println("Query execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"resume":    resume,
			"bio":       bio,
			"followers": followersCount,
			"following": followingCount,
		})
	}
}
