package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetUserFeed(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEmail := c.Query("email")
		log.Println(userEmail)

		// Step 1: Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", userEmail).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		log.Println(currentUserId)
		// Step 1

		// Step 2: Get the list of user IDs the current user is following
		var followedUserIds []int
		rows, err := db.Query("SELECT followed_user_id FROM follows WHERE following_user_id = $1", currentUserId)
		log.Println(rows)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		for rows.Next() {
			var followedUserId int
			if err := rows.Scan(&followedUserId); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			followedUserIds = append(followedUserIds, followedUserId)
		}
		rows.Close()
		log.Println(followedUserIds)
		// Step 2

		// Check if there are any followed user IDs
		var posts []models.Post
		if len(followedUserIds) == 0 {
			return c.JSON(posts) // or handle the empty result however you'd like
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Step 3: Build the query using Squirrel
		query := sb.Select("id", "user_id", "s3_key", "title", "game", "body", "status", "likes", "created_at").
			From("posts").
			Where(sq.Eq{"user_id": followedUserIds}).
			OrderBy("created_at DESC").
			Limit(10)

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		log.Println(sqlQuery)
		log.Println(args)
		if err != nil {
			log.Println("Prepare query 3")
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err = db.Query(sqlQuery, args...)
		if err != nil {
			log.Println("Execute query 3")
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		// Scan the results into a slice of Post structs
		for rows.Next() {
			var post models.Post
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.Title, &post.Game, &post.Body, &post.Status, &post.Likes, &post.CreatedAt); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			posts = append(posts, post)
		}
		// Step 3

		// Convert the posts to JSON and write the response
		return c.JSON(posts)
	}
}
