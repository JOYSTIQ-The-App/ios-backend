package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetUserBio(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve the bio based on the provided username
		query := sb.Select("profiles.bio").
			From("profiles").
			Join("users ON profiles.user_id = users.id").
			Where(sq.Eq{"users.username": username}).
			Limit(1) // Since username is unique, we only expect one match

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		row := db.QueryRow(sqlQuery, args...)
		var bio string
		err = row.Scan(&bio)
		if err == sql.ErrNoRows {
			bio = "" // Set blank string if no bio is found
		} else if err != nil {
			log.Println("Query execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"bio": bio,
		})
	}
}
