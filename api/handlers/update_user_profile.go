package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type ProfileUpdateRequest struct {
	Username string `json:"username"`
	Bio      string `json:"bio,omitempty"`
	Resume   string `json:"resume,omitempty"`
}

func UpdateGamerProfile(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(ProfileUpdateRequest)

		// Parse the body into the ProfileUpdateRequest struct
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Couldn't read request",
			})
		}

		// Ensure username was provided
		if request.Username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve the user's ID based on the provided username
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", request.Username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Attempt to insert a new profile with empty strings for the user, if it doesn't exist
		insertQuery := sb.Insert("profiles").Columns("user_id").
			Values(userID).
			Suffix("ON CONFLICT (user_id) DO NOTHING")

		insertSql, insertArgs, err := insertQuery.ToSql()
		if err != nil {
			log.Println("Prepare insert query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		_, err = db.Exec(insertSql, insertArgs...)
		if err != nil {
			log.Println("Insert operation failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Check if both bio and resume are empty
		if request.Bio == "" && request.Resume == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Both bio and resume are empty. No update was made.",
			})
		}

		// Prepare update operation, updating only the provided fields
		updateQuery := sb.Update("profiles")
		if request.Bio != "" {
			updateQuery = updateQuery.Set("bio", request.Bio)
		}
		if request.Resume != "" {
			updateQuery = updateQuery.Set("resume", request.Resume)
		}
		updateQuery = updateQuery.Where(sq.Eq{"user_id": userID})

		// Prepare and execute the update query
		updateSql, updateArgs, err := updateQuery.ToSql()
		if err != nil {
			log.Println("Prepare update query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		_, err = db.Exec(updateSql, updateArgs...)
		if err != nil {
			log.Println("Update operation failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Profile updated successfully.",
		})
	}
}
