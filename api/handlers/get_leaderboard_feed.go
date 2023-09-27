package handlers

import (
	"database/sql"
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

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the query using Squirrel
		query := sb.Select("p.id", "p.user_id", "p.s3_key", "p.media", "p.game", "p.body", "p.status", "p.created_at", "COUNT(lp.post_id) AS likes_count").
			From("posts p").
			Join("liked_posts lp ON p.id = lp.post_id").
			Where("lp.created_at >= ?", lastSunday).
			GroupBy("p.id", "p.user_id", "p.s3_key", "p.media", "p.game", "p.body", "p.status", "p.created_at").
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
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.Media, &post.Game, &post.Body, &post.Status, &post.CreatedAt, &post.LikesCount); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			posts = append(posts, post)
		}

		return c.JSON(posts)
	}
}
