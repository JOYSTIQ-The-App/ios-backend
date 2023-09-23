package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

// TO DO: Add Account. Currently it is defaulted to "public"
func CreateLike(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userName := c.Query("username")
		postId := c.Query("post_id")

		// Get user ID of the user with the given email
		var userId int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", userName).Scan(&userId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the INSERT query
		query := sb.Insert("liked_posts").Columns("user_id", "post_id").
			Values(userId, postId)

		// Pepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		_, err = db.Exec(sqlQuery, args...)
		if err != nil {
			log.Println("Execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"status": "success",
		})
	}
}
