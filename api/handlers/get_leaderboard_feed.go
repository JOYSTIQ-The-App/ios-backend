package handlers

import (
	"database/sql"
	"log"
	"time"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetLeaderboardFeed(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Determine the timestamp for the most recent Sunday
		lastSunday := time.Now().Truncate(time.Hour * 24)
		for lastSunday.Weekday() != time.Sunday {
			lastSunday = lastSunday.AddDate(0, 0, -1)
		}

		userEmail := c.Query("email")

		// Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", userEmail).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrel
		query := sb.Select(
			"p.id", "p.user_id", "p.s3_key", "p.thumbnail_s3_key", "p.media", "p.game", "p.body", "p.status", "p.likes", "p.comments", "p.created_at", "COUNT(lp.post_id) AS likes_count",
			"CASE WHEN ul.post_id IS NOT NULL THEN TRUE ELSE FALSE END AS user_liked",
			"users.username",
			"avatars.avatar_s3_key").
			From("posts p").
			Join("liked_posts lp ON p.id = lp.post_id").
			LeftJoin("(SELECT post_id FROM liked_posts WHERE user_id = ?) AS ul ON p.id = ul.post_id", currentUserId).
			Join("users ON p.user_id = users.id").
			LeftJoin("avatars ON p.user_id = avatars.user_id").
			Where(sq.And{
				sq.GtOrEq{"lp.created_at": lastSunday},
				sq.Eq{"p.media": "video"},
			}).
			GroupBy("p.id", "p.user_id", "p.s3_key", "p.thumbnail_s3_key", "p.media", "p.game", "p.body", "p.status", "p.likes", "p.comments", "p.created_at", "ul.post_id", "users.username", "avatars.avatar_s3_key").
			OrderBy("likes_count DESC").
			Limit(10)

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		// Scan the results into a slice of Post structs
		posts := make([]models.LeaderboardPost, 0)
		for rows.Next() {
			var post models.LeaderboardPost
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.ThumbnailS3Key, &post.Media, &post.Game, &post.Body, &post.Status, &post.Likes, &post.Comments, &post.CreatedAt, &post.LikesCount, &post.UserLiked, &post.Username, &post.AvatarS3Key); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			posts = append(posts, post)
		}

		return c.JSON(posts)
	}
}
