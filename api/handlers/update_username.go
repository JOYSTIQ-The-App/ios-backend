package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/validators"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type UpdateUsernameRequestBody struct {
	Email       string `json:"email"`
	NewUsername string `json:"new_username"`
}

func UpdateUsername(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get the email and new username
		var requestBody UpdateUsernameRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Check if the email exists and get the user ID
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", requestBody.Email).Scan(&userID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusBadRequest).SendString("Email does not exist.")
			}
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user ID.")
		}

		// Check if the new username is valid
		if !validators.IsValidUsername(requestBody.NewUsername) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid username format.")
		}

		// Update username in the database using Squirrel
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		query := sb.Update("users").
			Set("username", requestBody.NewUsername).
			Where(sq.Eq{"id": userID})

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to build query.")
		}

		_, err = db.Exec(sql, args...)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update username.")
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Username updated successfully",
		})
	}
}
