package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type DeleteSocialRequest struct {
	Username   string `json:"username"`
	SocialType string `json:"social_type"`
}

func DeleteUserSocial(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get deletion details
		var requestBody DeleteSocialRequest
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Validate social type using parsers
		social, err := parsers.ParseSocialType(requestBody.SocialType)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid social type.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve the user's ID based on the provided username
		var userID int
		err = db.QueryRow("SELECT id FROM users WHERE username = $1", requestBody.Username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Delete the specified social for the user
		query := sb.Delete("socials").
			Where(sq.Eq{"user_id": userID, "social": social})

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare delete query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		_, err = db.Exec(sqlQuery, args...)
		if err != nil {
			log.Println("Delete operation failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).SendString("Social deleted successfully.")
	}
}
