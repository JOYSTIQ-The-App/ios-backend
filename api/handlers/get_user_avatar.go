package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUserAvatar(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// First, retrieve the user's ID based on the provided username
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user details.")
		}

		// Now, retrieve the s3_key for the obtained user ID from the avatars table
		var s3Key sql.NullString
		err = db.QueryRow("SELECT s3_key FROM avatars WHERE user_id = $1", userID).Scan(&s3Key)
		if err == sql.ErrNoRows || !s3Key.Valid {
			return c.JSON(fiber.Map{
				"s3_key": nil,
			})
		} else if err != nil {
			log.Println("Error retrieving avatar s3_key:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve avatar.")
		}

		return c.JSON(fiber.Map{
			"s3_key": s3Key.String,
		})
	}
}
