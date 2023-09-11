// app.Get("/api/handlers/leaderboard", handlers.GetLeaderboardPosts(db))
// package handlers

// import (
// 	"database/sql"
// 	"time"

// 	"github.com/Masterminds/squirrel"
// 	"github.com/gofiber/fiber/v2"
// )

// func GetLeaderboardPosts(db *sql.DB) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Determine the timestamp for the most recent Sunday
// 		lastSunday := time.Now().Truncate(time.Hour * 24)
// 		for lastSunday.Weekday() != time.Sunday {
// 			lastSunday = lastSunday.AddDate(0, 0, -1)
// 		}

// 		// Build the query using Squirrel
// 		query := squirrel.Select("p.id", "p.title", "p.body", "COUNT(pl.id) AS likes_count", "p.created_at").
// 			From("Posts p").
// 			LeftJoin("post_likes pl ON p.id = pl.post_id AND pl.created_at >= ?", lastSunday).
// 			GroupBy("p.id").
// 			OrderBy("likes_count DESC").
// 			Limit(10)

// 		// Prepare and execute the query
// 		sqlQuery, args, err := query.ToSql()
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}

// 		rows, err := db.Query(sqlQuery, args...)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		defer rows.Close()

// 		// Scan the results into a slice of Post structs
// 		var posts []Post
// 		for rows.Next() {
// 			var post Post
// 			if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Likes, &post.CreatedAt); err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error": err.Error(),
// 				})
// 			}
// 			posts = append(posts, post)
// 		}

// 		// Return the posts as JSON
// 		return c.JSON(posts)
// 	}
// }
