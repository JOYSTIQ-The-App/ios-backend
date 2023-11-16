package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetGlobalFeed(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Query("email")
		lastSeenCreatedAt := c.Query("last_seen_created_at")

		if email == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Email not provided.")
		}

		var userId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userId)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Error retrieving user ID:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrel
		query := sb.Select("p.id", "p.user_id", "p.s3_key", "p.thumbnail_s3_key", "p.media", "p.game", "p.body", "p.status", "p.likes", "p.comments", "p.created_at",
			"CASE WHEN lp.post_id IS NOT NULL THEN TRUE ELSE FALSE END as user_liked",
			"users.username",
			"avatars.avatar_s3_key").
			From("posts p").
			LeftJoin("liked_posts lp ON p.id = lp.post_id AND lp.user_id = ?", userId).
			Join("users ON p.user_id = users.id").
			LeftJoin("avatars ON p.user_id = avatars.user_id").
			OrderBy("p.created_at DESC").
			Limit(10)

		// Only add this WHERE condition if lastSeenCreatedAt is not empty
		if lastSeenCreatedAt != "" {
			query = query.Where(sq.Lt{"p.created_at": lastSeenCreatedAt})
		}

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
		posts := make([]models.FeedPost, 0)
		for rows.Next() {
			var post models.FeedPost
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.ThumbnailS3Key, &post.Media, &post.Game, &post.Body, &post.Status, &post.Likes, &post.Comments, &post.CreatedAt, &post.UserLiked, &post.Username, &post.AvatarS3Key); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			posts = append(posts, post)
		}

		return c.JSON(posts)
	}
}
