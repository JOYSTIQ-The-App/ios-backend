package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetPostComments(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the post_id from the request
		postID := c.Query("post_id")
		if postID == "" {
			return c.Status(fiber.StatusBadRequest).SendString("post_id is required")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrel
		query := sb.Select("id", "post_id", "user_id", "text", "created_at").
			From("comments").
			Where(sq.Eq{"post_id": postID}).
			OrderBy("created_at ASC")

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			log.Println("Execute query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		// Scan the results into a slice of Comment structs
		comments := make([]models.Comment, 0)
		for rows.Next() {
			var comment models.Comment
			if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Text, &comment.CreatedAt); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			comments = append(comments, comment)
		}

		return c.JSON(comments)
	}
}
