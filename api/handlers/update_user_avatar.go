package handlers

import (
	"database/sql"
	"log"
	"time"

	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type AvatarRequestBody struct {
	Username    string `json:"username"`
	S3Key       string `json:"s3_key"`
	Environment string `json:"environment"`
}

func UpdateUserAvatar(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get post details
		var requestBody AvatarRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Validate background using parsers
		environment, err := parsers.ParseEnvironmentType(requestBody.Environment)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid environment type.",
			})
		}

		// Confirm presence of required fields
		if requestBody.Username == "" || requestBody.S3Key == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username and s3key are required.")
		}

		// Get user ID based on provided username
		var userID int
		err = db.QueryRow("SELECT id FROM users WHERE username = $1", requestBody.Username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user details.")
		}

		// Check if an avatar entry already exists for this user
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		var existingS3Key string
		err = db.QueryRow("SELECT s3_key FROM avatars WHERE user_id = $1", userID).Scan(&existingS3Key)

		if err == sql.ErrNoRows {
			// No avatar exists for this user; insert a new one
			insertQuery := sb.Insert("avatars").
				Columns("user_id", "s3_key", "environment").
				Values(userID, requestBody.S3Key, environment)
			_, err = insertQuery.RunWith(db).Exec()
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to create avatar entry.")
			}
			return c.Status(fiber.StatusOK).SendString("Avatar created successfully.")
		} else if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to check for existing avatar.")
		}

		// Avatar already exists for this user; update it
		updateQuery := sb.Update("avatars").
			Set("s3_key", requestBody.S3Key).
			Set("environment", environment).
			Set("updated_at", time.Now()).
			Where(sq.Eq{"user_id": userID})

		_, err = updateQuery.RunWith(db).Exec()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update avatar.")
		}

		return c.Status(fiber.StatusOK).SendString("Avatar updated successfully.")
	}
}
