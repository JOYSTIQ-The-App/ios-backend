package handlers

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsername(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEmail := c.Query("email")
		userIDStr := c.Query("user_id")

		var username string

		// Decide the query based on input (email or userID)
		if userEmail != "" {
			err := db.QueryRow("SELECT username FROM users WHERE email = $1", userEmail).Scan(&username)

			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		} else if userIDStr != "" {
			userID, err := strconv.Atoi(userIDStr) // Convert string userID to int
			if err != nil {
				log.Println("Invalid userID format:", err)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid userID format")
			}
			err = db.QueryRow("SELECT username FROM users WHERE id = $1", userID).Scan(&username)

			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("You must provide either an email or a userID")
		}

		return c.JSON(username)
	}
}
