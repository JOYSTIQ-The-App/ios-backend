package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq" // This is for the specific error check. Ensure it's imported.
)

type FollowRequestBody struct {
	FollowerUsername string `json:"follower_username"` // User who wants to follow
	FollowedUsername string `json:"followed_username"` // User who is getting followed
}

func CreateFollow(db *sql.DB) fiber.Handler {
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

		// Attempt to insert the new follow relationship into the database
		query := sb.Insert("follows").
			Columns("following_user_id", "followed_user_id").
			Values(followerID, followedID)

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare insert query failed:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		_, err = db.Exec(sqlQuery, args...)
		if err != nil {
			// Check if the error is due to unique constraint violation
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" { // unique_violation
					return c.Status(fiber.StatusBadRequest).SendString("Follow relationship already exists.")
				}
			}
			log.Println("Insert operation failed:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).SendString("Followed successfully.")
	}
}
