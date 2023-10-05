package handlers

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetFollowersList(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// First, retrieve the user's ID based on the provided username
		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Retrieve the list of usernames who follow this user
		query := sb.Select("u.username").
			From("follows AS f").
			Join("users AS u ON f.following_user_id = u.id").
			Where(sq.Eq{"f.followed_user_id": userID})

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			log.Println("Execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		// Parse the list of follower usernames
		var followerUsernames []string
		for rows.Next() {
			var followerUsername string
			if err := rows.Scan(&followerUsername); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			followerUsernames = append(followerUsernames, followerUsername)
		}

		return c.JSON(fiber.Map{
			"followers": followerUsernames,
		})
	}
}
