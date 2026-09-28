package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")

		// Ensure username was provided
		if username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username not provided.")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Retrieve user's ID, avatar_s3_key, and environment
		query := sb.Select("users.id", "avatars.avatar_s3_key", "avatars.environment").
			From("users").
			LeftJoin("avatars ON users.id = avatars.user_id").
			Where(sq.Eq{"users.username": username}).
			Limit(1)

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare avatar query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		var userID int
		var avatarS3Key, environment sql.NullString // Using NullString to handle potential NULL values

		err = db.QueryRow(sqlQuery, args...).Scan(&userID, &avatarS3Key, &environment)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString("User not found.")
		} else if err != nil {
			log.Println("Avatar query execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the count of followers
		var followersCount int
		err = db.QueryRow("SELECT COUNT(*) FROM follows WHERE followed_user_id = $1", userID).Scan(&followersCount)
		if err != nil {
			log.Println("Error counting followers", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the count of users the user is following
		var followingCount int
		err = db.QueryRow("SELECT COUNT(*) FROM follows WHERE following_user_id = $1", userID).Scan(&followingCount)
		if err != nil {
			log.Println("Error counting followed users", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Retrieve the socials for the obtained user ID
		query = sb.Select("social", "social_username").
			From("socials").
			Where(sq.Eq{"user_id": userID})

		sqlQuery, args, err = query.ToSql()
		if err != nil {
			log.Println("Prepare socials query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rows, err := db.Query(sqlQuery, args...)
		if err != nil {
			log.Println("Execute socials query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		userSocials := make(map[string]string, 0)
		for rows.Next() {
			var socialType, socialUsername string
			if err := rows.Scan(&socialType, &socialUsername); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			userSocials[socialType] = socialUsername
		}

		// Now, retrieve the resume and bio for the obtained user ID
		query = sb.Select("profiles.resume", "profiles.bio").From("profiles").
			Where(sq.Eq{"user_id": userID}).
			Limit(1) // Since user ID is unique, we only expect one match

		sqlQuery, args, err = query.ToSql()
		if err != nil {
			log.Println("Prepare profile query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		row := db.QueryRow(sqlQuery, args...)
		var resume, bio string
		err = row.Scan(&resume, &bio)
		if err == sql.ErrNoRows {
			resume, bio = "", "" // Set blank strings if no profile is found
		} else if err != nil {
			log.Println("Profile query execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Retrieve the posts for the obtained user ID
		query = sb.Select(
			"posts.id", "posts.user_id", "s3_key", "thumbnail_s3_key", "media", "game", "body", "status", "likes", "comments", "posts.created_at",
			"CASE WHEN liked_posts.post_id IS NOT NULL THEN TRUE ELSE FALSE END as user_liked",
			"users.username",
			"avatars.avatar_s3_key").
			From("posts").
			LeftJoin("liked_posts ON posts.id = liked_posts.post_id AND liked_posts.user_id = ?", userID).
			Join("users ON posts.user_id = users.id").
			LeftJoin("avatars ON posts.user_id = avatars.user_id").
			Where(sq.Eq{"posts.user_id": userID}).
			OrderBy("created_at DESC")

		sqlQuery, args, err = query.ToSql()
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
		posts := make([]models.FeedPost, 0)
		for rows.Next() {
			var post models.FeedPost
			if err := rows.Scan(&post.ID, &post.UserID, &post.S3Key, &post.ThumbnailS3Key, &post.Media, &post.Game, &post.Body, &post.Status, &post.Likes, &post.Comments, &post.CreatedAt, &post.UserLiked, &post.Username, &post.AvatarS3Key); err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			posts = append(posts, post)
		}

		return c.JSON(fiber.Map{
			"avatar_s3_key": avatarS3Key,
			"bio":           bio,
			"environment":   environment.String,
			"followers":     followersCount,
			"following":     followingCount,
			"posts":         posts,
			"resume":        resume,
			"socials":       userSocials,
		})
	}
}
