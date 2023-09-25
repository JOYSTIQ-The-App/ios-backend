package handlers

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func SearchUsernames(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		searchTerm := c.Query("username") // Assuming the query parameter is named "q"

		// If search term is empty, return a bad request
		if searchTerm == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Search term cannot be empty.",
			})
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrel
		query := sb.Select("username").
			From("users").
			Where("username LIKE ?", "%"+searchTerm+"%").
			Limit(10) // Limit to top 10 results, adjust as needed

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		// Scan the results into a slice of strings (usernames)
		usernames := make([]string, 0)
		for rows.Next() {
			var username string
			if err := rows.Scan(&username); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			usernames = append(usernames, username)
		}

		return c.JSON(usernames)
	}
}
