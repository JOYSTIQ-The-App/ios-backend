package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetGlobalFeed(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Build the query using Squirrel
		query := squirrel.Select("id", "user_id", "s3_key", "title", "game", "body", "status", "likes", "created_at").
			From("Posts").
			OrderBy("created_at DESC").
			Limit(10)

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		// Scan the results into a slice of Post structs
		var posts []models.Post
		for rows.Next() {
			var post models.Post
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.Title, &post.Game, &post.Body, &post.Status, &post.Likes, &post.CreatedAt); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			posts = append(posts, post)
		}

		// Convert the posts to JSON and write the response
		return c.JSON(posts)
	}
}
