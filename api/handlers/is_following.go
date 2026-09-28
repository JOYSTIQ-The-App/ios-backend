package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func IsFollowing(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		follower := c.Query("username")
		followed := c.Query("following_username")

		// Ensure both usernames were provided
		if follower == "" || followed == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Both follower and followed usernames must be provided.")
		}

		// Retrieve user IDs for both usernames
		var followerID, followedID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", follower).Scan(&followerID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("Follower user not found.")
		} else if err != nil {
			log.Println("Error retrieving follower ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		err = db.QueryRow("SELECT id FROM users WHERE username = $1", followed).Scan(&followedID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("Followed user not found.")
		} else if err != nil {
			log.Println("Error retrieving followed ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Check if follower is following followed user
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM follows WHERE following_user_id = $1 AND followed_user_id = $2)", followerID, followedID).Scan(&exists)
		if err != nil {
			log.Println("Error checking following status", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Return the result
		return c.JSON(fiber.Map{
			"isFollowing": exists,
		})
	}
}
