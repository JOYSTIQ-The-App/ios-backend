package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	"github.com/JOYSTIQ-The-App/Backend/api/services"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type PostRequestBody struct {
	S3Key  sql.NullString `json:"s3_key"`
	Media  string         `json:"media"`
	Title  string         `json:"title"`
	Game   string         `json:"game"`
	Body   string         `json:"body"`
	Status string         `json:"status"`
}

func CreatePost(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEmail := c.Query("email")

		// Validation
		if userEmail == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Email is required.")
		}

		// Get user ID of the user with the given email
		var currentUserId int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", userEmail).Scan(&currentUserId)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to retrieve user.",
			})
		}

		// Parse body to get post details
		var requestBody PostRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to parse request body.",
			})
		}

		// Ensure there is no profanity
		detector := services.CustomProfanityDetector()
		if detector.IsProfane(requestBody.Body) {
			requestBody.Body = detector.Censor(requestBody.Body)
		}

		// Validate media and status using parsers
		media, err := parsers.ParseMediaType(requestBody.Media)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid media type.",
			})
		}

		status, err := parsers.ParseStatusType(requestBody.Status)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid status type.",
			})
		}

		// Construct the Post model
		post := models.Post{
			UserID: currentUserId,
			S3Key:  requestBody.S3Key,
			Media:  media,
			Game:   requestBody.Game,
			Body:   requestBody.Body,
			Status: status,
		}

		// Insert new post into the database using Squirrel
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		query := sb.Insert("posts").
			Columns("user_id", "s3_key", "media", "game", "body", "status").
			Values(post.UserID, post.S3Key, post.Media, post.Game, post.Body, post.Status)

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to build query.",
			})
		}

		_, err = db.Exec(sql, args...)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create post.",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Post created successfully.",
		})
	}
}
