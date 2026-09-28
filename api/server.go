package main

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/handlers"
	"github.com/JOYSTIQ-The-App/Backend/api/jobs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

func main() {
	connStr := "user=postgres password=postgres host=localhost port=5432 dbname=joystiq sslmode=disable"
	// connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	// Update Comments Cron Job
	updateCommentsCron := cron.New(cron.WithSeconds())
	_, err = updateCommentsCron.AddFunc("*/10 * * * * *", func() {
		err := jobs.UpdateCommentCounts(db)
		if err != nil {
			log.Println("Error running UpdateCommentCounts:", err)
		}
	})

	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	updateCommentsCron.Start()

	// Update Comments Cron Job
	updateLikesCron := cron.New(cron.WithSeconds())
	_, err = updateLikesCron.AddFunc("*/10 * * * * *", func() {
		err := jobs.UpdateLikesCount(db)
		if err != nil {
			log.Println("Error running UpdateLikesCount:", err)
		}
	})

	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	updateLikesCron.Start()

	// Fiber
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Delete("handlers/delete_like", handlers.DeleteLike(db))
	app.Delete("handlers/delete_post", handlers.DeletePost(db))
	app.Delete("handlers/delete_user_social", handlers.DeleteUserSocial(db))
	app.Delete("handlers/delete_follow", handlers.DeleteFollow(db))

	app.Get("handlers/get_global_feed", handlers.GetGlobalFeed(db))
	app.Get("handlers/get_leaderboard_feed", handlers.GetLeaderboardFeed(db))
	app.Get("handlers/get_user_feed", handlers.GetUserFeed(db))
	app.Get("handlers/get_username", handlers.GetUsername(db))
	app.Get("handlers/search_usernames", handlers.SearchUsernames(db))
	app.Get("handlers/get_post_comments", handlers.GetPostComments(db))
	app.Get("handlers/get_user_profile", handlers.GetUserProfile(db))
	app.Get("handlers/get_user_posts", handlers.GetUserPosts(db))
	app.Get("handlers/get_followers_list", handlers.GetFollowersList(db))
	app.Get("handlers/get_following_list", handlers.GetFollowingList(db))
	app.Get("handlers/get_user_avatar", handlers.GetUserAvatar(db))
	app.Get("handlers/get_user_socials", handlers.GetUserSocials(db))
	app.Get("/handlers/check_username_availability", handlers.CheckUsernameAvailability(db))
	app.Get("/handlers/is_email_registered", handlers.IsEmailRegistered(db))
	app.Get("handlers/is_following", handlers.IsFollowing(db))
	app.Get("handlers/get_user_bio", handlers.GetUserBio(db))
	app.Get("handlers/get_gamer_profile", handlers.GetGamerProfile(db))

	app.Post("handlers/create_like", handlers.CreateLike(db))
	app.Post("handlers/create_post", handlers.CreatePost(db))
	app.Post("handlers/create_user", handlers.CreateUser(db))
	app.Post("handlers/create_comment", handlers.CreateComment(db))
	app.Post("handlers/create_follow", handlers.CreateFollow(db))

	app.Patch("handlers/update_user_profile", handlers.UpdateGamerProfile(db))
	app.Patch("handlers/update_user_post", handlers.UpdateUserPost(db))
	app.Patch("handlers/update_user_avatar", handlers.UpdateUserAvatar(db))
	app.Patch("handlers/update_user_socials", handlers.UpdateUserSocials(db))
	app.Patch("handlers/update_username", handlers.UpdateUsername(db))
	app.Patch("handlers/report_post", handlers.ReportPost(db))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	// log.Fatal(app.Listen("0.0.0.0:8080"))
	log.Fatal(app.Listen(":8080"))
}
