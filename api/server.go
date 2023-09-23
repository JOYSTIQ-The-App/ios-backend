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

	app.Get("handlers/get_global_feed", handlers.GetGlobalFeed(db))
	app.Get("handlers/get_leaderboard_feed", handlers.GetLeaderboardFeed(db))
	app.Get("handlers/get_user_feed", handlers.GetUserFeed(db))
	app.Get("handlers/get_username", handlers.GetUsername(db))

	app.Post("handlers/create_like", handlers.CreateLike(db))
	app.Post("handlers/create_post", handlers.CreatePost(db))
	app.Post("handlers/create_user", handlers.CreateUser(db))

	log.Fatal(app.Listen(":8080"))
}
