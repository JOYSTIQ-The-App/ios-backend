package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func DeleteLike(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userName := c.Query("username")
		postId := c.Query("post_id")

		// Get user ID of the user with the given username
		var userId int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", userName).Scan(&userId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the DELETE query
		query := sb.Delete("liked_posts").
			Where(sq.Eq{"user_id": userId}).
			Where(sq.Eq{"post_id": postId})

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare delete query failed", err)
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
