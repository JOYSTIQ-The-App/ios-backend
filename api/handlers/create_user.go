package handlers

import (
	"database/sql"
	"log"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
	"github.com/JOYSTIQ-The-App/Backend/api/parsers"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

type UserRequestBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Account  string `json:"account"`
}

// TO DO: Add Account. Currently it is defaulted to "public"
func CreateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body to get user details
		var requestBody UserRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body.")
		}

		// Validate account and role using parsers
		account, err := parsers.ParseAccountType(requestBody.Account)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid account type.")
		}

		role, err := parsers.ParseRoleType(requestBody.Role)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid role type.")
		}

		// Construct the Post model
		user := models.User{
			Email:    requestBody.Email,
			Username: requestBody.Username,
			Role:     role,
			Account:  account,
		}

		// Insert new post into the database using Squirrel
		sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		query := sb.Insert("users").
			Columns("email", "username", "role", "account").
			Values(user.Email, user.Username, user.Role, user.Account).
			Suffix("RETURNING id")

		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to build query.")
		}

		err = db.QueryRow(sql, args...).Scan(&user.ID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user.")
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"user_id": user.ID,
		})
	}
}
