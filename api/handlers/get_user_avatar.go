package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetUserAvatar(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve the user's ID based on the provided username
		query := sb.Select("id").From("users").Where(sq.Eq{"username": username}).Limit(1)
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Error preparing query to get user ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user details.")
		}

		var userID int
		err = db.QueryRow(sqlQuery, args...).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user details.")
		}

		// Retrieve the s3_key and environment for the obtained user ID from the avatars table
		query = sb.Select("s3_key", "environment").From("avatars").Where(sq.Eq{"user_id": userID}).Limit(1)
		sqlQuery, args, err = query.ToSql()
		if err != nil {
			log.Println("Error preparing query to get avatar details:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve avatar details.")
		}

		var s3Key, environment sql.NullString
		err = db.QueryRow(sqlQuery, args...).Scan(&s3Key, &environment)
		if err == sql.ErrNoRows {
			return c.JSON(fiber.Map{
				"s3_key":      nil,
				"environment": "classic",
			})
		} else if err != nil {
			log.Println("Error retrieving avatar s3_key and environment:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve avatar details.")
		}

		return c.JSON(fiber.Map{
			"s3_key":      s3Key.String,
			"environment": environment.String,
		})
	}
}
