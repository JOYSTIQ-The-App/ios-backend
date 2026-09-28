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

// Mock list of reserved usernames
// var validGames = []string{
// 	"VALORANT",
// 	"Call of Duty",
// 	"Call of Duty: Warzone",
// 	"Apex Legends",
// 	"Overwatch",
// 	"Battlefield",
// 	"Halo",
// 	"Rust",
// 	"Escape from Tarkov",
// 	"Fortnite",
// 	"League of Legends",
// 	"Among us",
// 	"Sea of Thieves",
// 	"Stardew Valley",
// 	"Mortal Combat 1",
// 	"Fall guys",
// 	"Rocket league",
// 	"Grand Theft auto",
// 	"Clash royale",
// 	"Minecraft",
// 	"CS:GO",
// 	"Counter-Strike 2",
// 	"Super Smash bros",
// 	"Bloodhunt",
// 	"Multiversus",
// 	"God of War",
// 	"Spiderman",
// 	"Red Dead Redemption",
// 	"Resident Evil",
// 	"Dota 2",
// 	"Lost Ark",
// 	"Teamfight Tactics",
// 	"Dead by Daylight",
// 	"Rainbow Six Siege",
// 	"Genshin Impact",
// 	"Starfield",
// 	"PlayerUnknown’s Battlegrounds",
// 	"EA Sports FC 24",
// 	"NBA 2k24",
// 	"Hearthstone",
// 	"Path of Exile",
// 	"Dark and Darker",
// 	"Street Fighter 6",
// 	"Madden NFL 24",
// 	"Subnautica",
// 	"Phasmophobia",
// 	"Elden Ring",
// 	"DayZ",
// }

// // Function to check if a given username is reserved
// func isValidGame(game string) bool {
// 	for _, validGame := range validGames {
// 		if strings.EqualFold(game, validGame) { // Case-insensitive check
// 			return true
// 		}
// 	}
// 	return false
// }

type PostRequestBody struct {
	S3Key          sql.NullString `json:"s3_key"`
	ThumbnailS3Key sql.NullString `json:"thumbnail_s3_key"`
	Media          string         `json:"media"`
	Title          string         `json:"title"`
	Game           string         `json:"game"`
	Body           string         `json:"body"`
	Status         string         `json:"status"`
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

		if requestBody.Game != "" {
			if detector.IsProfane(requestBody.Game) {
				requestBody.Game = ""
			}
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
			UserID:         currentUserId,
			S3Key:          requestBody.S3Key,
			ThumbnailS3Key: requestBody.ThumbnailS3Key,
			Media:          media,
			Game:           requestBody.Game,
			Body:           requestBody.Body,
			Status:         status,
		}

		// Insert new post into the database using Squirrel
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		query := sb.Insert("posts").
			Columns("user_id", "s3_key", "thumbnail_s3_key", "media", "game", "body", "status").
			Values(post.UserID, post.S3Key, post.ThumbnailS3Key, post.Media, post.Game, post.Body, post.Status)

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
