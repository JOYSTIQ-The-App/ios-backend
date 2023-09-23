package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

// TO DO: Add Account. Currently it is defaulted to "public"
func CreateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEmail := c.Query("email")
		userName := c.Query("username")

		roleStr := c.Query("role")
		userRole, err := parsers.ParseRoleType(roleStr)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid role type provided")
		}

		// Create the Squirrel statement builder with PostgreSQL dollar format
		var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		// Build the INSERT query
		query := sb.Insert("users").Columns("email", "username", "role").
			Values(userEmail, userName, userRole).
			Suffix("RETURNING id")

		// Pepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		var userId int
		if err := db.QueryRow(sqlQuery, args...).Scan(&userId); err != nil {
			log.Println("Execution failed", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"user_id": userId,
		})
	}
}
