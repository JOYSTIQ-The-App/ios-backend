package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckUsernameAvailability(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the username from the request
		username := c.Query("username")

		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username must be provided.",
			})
		}

		// Check if the username exists in the database
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check username availability.",
			})
		}

		// If the username exists, return false, else true
		isAvailable := !exists

		return c.JSON(fiber.Map{
			"availability": isAvailable,
		})
	}
}
