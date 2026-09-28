package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func DeletePost(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEmail := c.Query("email")
		postID := c.Query("post_id")

		// Ensure postID was provided
		if postID == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Post ID not provided.")
		}

		// Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", userEmail).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Confirm if the post belongs to the user with the provided email
		var postUserID int
		err = db.QueryRow("SELECT user_id FROM posts WHERE id = $1", postID).Scan(&postUserID)
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

		// Start transaction
		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to start transaction.")
		}

		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Delete likes from "liked_posts" associated with the post
		likesQuery := sb.Delete("liked_posts").Where(sq.Eq{"post_id": postID})
		likesSql, likesArgs, err := likesQuery.ToSql()
		if err != nil {
			return rollbackOnError(tx, err)
		}
		_, err = tx.Exec(likesSql, likesArgs...)
		if err != nil {
			return rollbackOnError(tx, err)
		}

		// Delete comments associated with the post
		commentsQuery := sb.Delete("comments").Where(sq.Eq{"post_id": postID})
		commentsSql, commentsArgs, err := commentsQuery.ToSql()
		if err != nil {
			return rollbackOnError(tx, err)
		}
		_, err = tx.Exec(commentsSql, commentsArgs...)
		if err != nil {
			return rollbackOnError(tx, err)
		}

		// Delete the post itself
		postQuery := sb.Delete("posts").Where(sq.Eq{"id": postID})
		postSql, postArgs, err := postQuery.ToSql()
		if err != nil {
			return rollbackOnError(tx, err)
		}
		_, err = tx.Exec(postSql, postArgs...)
		if err != nil {
			return rollbackOnError(tx, err)
		}

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to commit transaction.")
		}

		return c.Status(fiber.StatusOK).SendString("Post, and associated likes and comments, deleted successfully.")
	}
}

// Helper function to handle rollback on errors
func rollbackOnError(tx *sql.Tx, err error) error {
	log.Println(err)
	tx.Rollback()
	return fiber.NewError(fiber.StatusInternalServerError, "Failed to process request. Transaction rollback initiated.")
}
