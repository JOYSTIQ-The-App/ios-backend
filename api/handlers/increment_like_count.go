package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func IncrementLikeCount(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postId := c.Query("id") // Assuming the ID of the post is passed as a URL parameter

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the UPDATE query to increment the likes count
		query := sb.Update("posts").Set("likes", sq.Expr("likes + 1")).Where(sq.Eq{"id": postId})

		// Prepare and execute the query
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
