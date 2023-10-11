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
		query := sb.Select("c.id", "c.post_id", "c.user_id", "u.username", "c.text", "c.created_at").
			From("comments c").
			Join("users u ON c.user_id = u.id").
			Where(sq.Eq{"c.post_id": postID}).
			OrderBy("c.created_at ASC")

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

		// Define a struct to hold the comment data including the username
		type CommentWithUsername struct {
			models.Comment
			Username string `json:"username"`
		}

		comments := make([]CommentWithUsername, 0)
		for rows.Next() {
			var comment CommentWithUsername
			if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Username, &comment.Text, &comment.CreatedAt); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			comments = append(comments, comment)
		}

		return c.JSON(comments)
	}
}
