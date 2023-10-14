package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func DeleteFollow(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody FollowRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Ensure both usernames were provided
		if requestBody.FollowerUsername == "" || requestBody.FollowedUsername == "" {
			return c.Status(fiber.StatusBadRequest).SendString("One of the usernames was blank")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve the user IDs for the follower and the followed user
		var followerID, followedID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", requestBody.FollowerUsername).Scan(&followerID)
		if err != nil {
			log.Println("Error retrieving follower ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		err = db.QueryRow("SELECT id FROM users WHERE username = $1", requestBody.FollowedUsername).Scan(&followedID)
		if err != nil {
			log.Println("Error retrieving followed user ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Delete the follow relationship
		query := sb.Delete("follows").
			Where(sq.Eq{"following_user_id": followerID, "followed_user_id": followedID})

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare delete query failed:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		result, err := db.Exec(sqlQuery, args...)
		if err != nil {
			log.Println("Delete operation failed:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println("Failed to get affected rows:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to determine affected rows.")
		}

		// If no rows were deleted, inform the user that the follow relationship was not found
		if rowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).SendString("Follow relationship not found.")
		}

		return c.Status(fiber.StatusOK).SendString("Unfollowed successfully.")
	}
}
