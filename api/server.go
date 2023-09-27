package main

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/handlers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres host=localhost port=5432 dbname=joystiq sslmode=disable" // Adjust connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Delete("handlers/delete_like", handlers.DeleteLike(db))

	app.Get("handlers/get_global_feed", handlers.GetGlobalFeed(db))
	app.Get("handlers/get_leaderboard_feed", handlers.GetLeaderboardFeed(db))
	app.Get("handlers/get_user_feed", handlers.GetUserFeed(db))
	app.Get("handlers/get_username", handlers.GetUsername(db))
	app.Get("handlers/search_usernames", handlers.SearchUsernames(db))
	app.Get("handlers/get_post_comments", handlers.GetPostComments(db))

	app.Post("handlers/create_like", handlers.CreateLike(db))
	app.Post("handlers/create_post", handlers.CreatePost(db))
	app.Post("handlers/create_user", handlers.CreateUser(db))
	app.Post("handlers/increment_like_count", handlers.IncrementLikeCount(db))
	app.Post("handlers/decrement_like_count", handlers.DecrementLikeCount(db))
	app.Post("handlers/create_comment", handlers.CreateComment(db))

	app.Patch("handlers/update_likes_count", handlers.UpdateLikesCount(db))
	app.Patch("handlers/update_comments_count", handlers.UpdateCommentCounts(db))

	// // Likes
	// app.Delete("/likes", handlers.DeleteLike(db))
	// app.Post("/likes", handlers.CreateLike(db))
	// app.Post("/likes/increment", handlers.IncrementLikeCount(db))
	// app.Post("/likes/decrement", handlers.DecrementLikeCount(db))

	// // Feed
	// app.Get("/feeds/global", handlers.GetGlobalFeed(db))
	// app.Get("/feeds/leaderboard", handlers.GetLeaderboardFeed(db))
	// app.Get("/feeds/user", handlers.GetUserFeed(db))

	// // Users
	// app.Get("/users/username", handlers.GetUsername(db))
	// app.Get("/users/search", handlers.SearchUsernames(db))
	// app.Post("/users", handlers.CreateUser(db))

	// // Posts
	// app.Post("/posts", handlers.CreatePost(db))
	// app.Get("/posts/comments", handlers.GetPostComments(db))

	// // Comments
	// app.Post("/comments", handlers.CreateComment(db))

	// Initialize the cron package
	// c := cron.New(cron.WithSeconds())
	// _, err = c.AddFunc("0 0 * * * *", func() { // This runs every hour at minute 0. Adjust as necessary.
	//     log.Println("Running the UpdateLikesCount job...")
	//     handler := handlers.UpdateLikesCount(db)
	//     // Since this handler expects a fiber context (which isn't available in this cron function),
	//     // you might want to refactor your handler function to work without the context when called from cron.
	//     handler(nil)
	// })

	// if err != nil {
	//     log.Fatalf("Error scheduling cron job: %v", err)
	// }

	// c.Start()

	log.Fatal(app.Listen(":8080"))
}
