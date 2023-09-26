package handlers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type CommentRequestBody struct {
	PostID   int    `json:"post_id"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func CreateComment(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get comment details
		var requestBody CommentRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Validate comment text
		trimmedText := strings.TrimSpace(requestBody.Text)
		if len(trimmedText) == 0 || len(trimmedText) > 500 {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid comment text length.")
		}

		// Fetch UserID based on the provided Username
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username=$1", requestBody.Username).Scan(&userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid username.")
		}

		// Check if post_id exists
		var postExists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id=$1)", requestBody.PostID).Scan(&postExists)
		if err != nil || !postExists {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid post_id.")
		}

		// Construct the Comment model
		comment := models.Comment{
			PostID: requestBody.PostID,
			UserID: userID,
			Text:   trimmedText,
		}

		// Insert new comment into the database using Squirrel
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		query := sb.Insert("comments").
			Columns("post_id", "user_id", "text").
			Values(comment.PostID, comment.UserID, comment.Text).
			Suffix("RETURNING id, created_at, updated_at")

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to build query.")
		}

		err = db.QueryRow(sql, args...).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create comment.")
		}

		return c.JSON(comment)
	}
}
