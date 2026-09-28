package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type ReportPostRequestBody struct {
	Username string `json:"username"`
	PostID   int    `json:"post_id"`
}

func ReportPost(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var requestBody ReportPostRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Get user ID of the user with the given username
		var userId int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", requestBody.Username).Scan(&userId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Update the post status to 'reported'
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		query := sb.Update("posts").
			Set("status", "reported").
			Where(sq.Eq{"id": requestBody.PostID})

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to build update query.")
		}

		result, err := db.Exec(sql, args...)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to report post.")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve affected rows.")
		}
		if rowsAffected == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid post ID. No post was reported.")
		}

		// Log the action
		log.Println(requestBody.Username, "reported post:", requestBody.PostID)

		return c.Status(fiber.StatusOK).SendString("Post reported successfully.")
	}
}
