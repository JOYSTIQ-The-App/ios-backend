package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func IsEmailRegistered(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve email from query parameter
		email := c.Query("email")
		if email == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Email is required.")
		}

		// Check if the email exists in the database
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check email registration status.",
			})
		}

		return c.JSON(fiber.Map{
			"isRegistered": exists,
		})
	}
}
