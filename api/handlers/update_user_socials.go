package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type UpdateSocialsRequest struct {
	Username       string `json:"username"`
	SocialType     string `json:"social_type"`
	SocialUsername string `json:"social_username"`
}

func UpdateUserSocials(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get socials details
		var requestBody UpdateSocialsRequest
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

		// Check if the social already exists for the user
		var existingSocial string
		err = db.QueryRow("SELECT social FROM socials WHERE user_id = $1 AND social = $2", userID, social).Scan(&existingSocial)

		if err != nil && err != sql.ErrNoRows {
			log.Println("Error checking existing social", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if existingSocial == "" {
			// Insert new social
			query := sb.Insert("socials").
				Columns("user_id", "social", "social_username").
				Values(userID, social, requestBody.SocialUsername)
			sqlQuery, args, err := query.ToSql()
			if err != nil {
				log.Println("Prepare insert query failed", err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			_, err = db.Exec(sqlQuery, args...)
			if err != nil {
				log.Println("Insert operation failed", err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		} else {
			// Update existing social username
			query := sb.Update("socials").
				Set("social_username", requestBody.SocialUsername).
				Where(sq.Eq{"user_id": userID, "social": social})
			sqlQuery, args, err := query.ToSql()
			if err != nil {
				log.Println("Prepare update query failed", err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			_, err = db.Exec(sqlQuery, args...)
			if err != nil {
				log.Println("Update operation failed", err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}

		return c.Status(fiber.StatusOK).SendString("Social updated successfully.")
	}
}
