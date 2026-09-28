package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type UpdatePostRequestBody struct {
	Email    string `json:"email"`
	PostID   int    `json:"post_id"`
	BodyText string `json:"body_text"`
}

func UpdateUserPost(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var requestBody UpdatePostRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", requestBody.Email).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Confirm if the post belongs to the user with the provided email
		var postUserID int
		err = db.QueryRow("SELECT user_id FROM posts WHERE id = $1", requestBody.PostID).Scan(&postUserID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("Post not found.")
		} else if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve post details.")
		}

		// If the post does not belong to the user, return forbidden
		if postUserID != currentUserId {
			return c.Status(fiber.StatusForbidden).SendString("Permission denied. This post does not belong to you.")
		}

		// Update the post body
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		query := sb.Update("posts").
			Set("body", requestBody.BodyText).
			Where(sq.Eq{"id": requestBody.PostID})

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to build update query.")
		}

		_, err = db.Exec(sql, args...)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update post.")
		}

		return c.Status(fiber.StatusOK).SendString("Post updated successfully.")
	}
}
