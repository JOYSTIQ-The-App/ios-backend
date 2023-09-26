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

		// Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", userEmail).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the list of user IDs the current user is following
		var followedUserIds []int
		rows, err := db.Query("SELECT followed_user_id FROM follows WHERE following_user_id = $1", currentUserId)
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

		// Check if there are any followed user IDs
		var posts []models.UserFeedPost
		if len(followedUserIds) == 0 {
			return c.JSON(posts)
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrels
		query := sb.Select(
			"posts.id", "posts.user_id", "s3_key", "media", "game", "body", "status", "likes", "comments", "posts.created_at",
			"CASE WHEN liked_posts.post_id IS NOT NULL THEN TRUE ELSE FALSE END as user_liked").
			From("posts").
			LeftJoin("liked_posts ON posts.id = liked_posts.post_id AND liked_posts.user_id = ?", currentUserId).
			Where(sq.Eq{"posts.user_id": followedUserIds}).
			OrderBy("created_at DESC").
			Limit(10)

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err = db.Query(sqlQuery, args...)
		if err != nil {
			log.Println("Execute query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		// Scan the results into a slice of Post structs
		for rows.Next() {
			var post models.UserFeedPost
			var userLiked bool
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.Media, &post.Game, &post.Body, &post.Status, &post.Likes, &post.Comments, &post.CreatedAt, &userLiked); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			post.UserLiked = userLiked // Assuming you have added this field to the Post model
			posts = append(posts, post)
		}

		return c.JSON(posts)
	}
}
