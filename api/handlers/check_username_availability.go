package handlers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/JOYSTIQ-The-App/Backend/api/services"
	"github.com/JOYSTIQ-The-App/Backend/api/validators"
	"github.com/gofiber/fiber/v2"
)

// Mock list of reserved usernames
var reservedUsernames = []string{
	"admin",
	"administrator",
	"guest",
	"support",
	"help",
	"root",
	"user",
	"staff",
	"official",
	"moderator",
	"mod",
	"xQc",
	"Jynxzi",
	"Loltyler1",
	"KaiCenat",
	"FanumTV",
	"IShowSpeed",
	"Pewdiepie",
	"Markiplier",
	"Ali-A",
	"Lachlan",
	"Jacksepticeye",
	"VanossGaming",
	"Ninja",
	"Nickmercs",
	"Timthetatman",
	"DrDisrespect",
	"Cloakzy",
	"Clix",
	"Pokimane",
	"Valkyrae",
	"Dakotaz",
	"Tfue",
	"DrLupo",
	"SypherPK",
	"Bugha",
	"Subroza",
	"Shanks",
	"Tarik",
	"TenZ",
	"Shroud",
	"S1mple",
	"Summit1g",
	"Aceu",
	"Faide",
	"ImperialHal",
	"LuLu",
	"lululuvely",
	"Apical",
	"Nooch",
	"Cotts",
	"OPTIC_H3CZ",
	"SCUMP",
	"ALEX",
	"ANDREA",
	"AYYJAYY",
	"BIGTYMER",
	"BLAKE",
	"DI3SEL",
	"FORMAL",
	"HITCHARIIDE",
	"HOLLY",
	"JORGE",
	"KARMA",
	"KENNY",
	"LUCID",
	"MANIAC",
	"MBOZE",
	"METHODZ",
	"PRED",
	"RETALS",
	"SEBAS",
	"SHOTTZY",
	"TEEP",
	"ZLANER",
}

// Function to check if a given username is reserved
func isUsernameReserved(username string) bool {
	for _, reservedUsername := range reservedUsernames {
		if strings.EqualFold(username, reservedUsername) { // Case-insensitive check
			return true
		}
	}
	return false
}

func CheckUsernameAvailability(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the username from the request
		username := c.Query("username")

		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username must be provided.",
			})
		}

		// Check if the new username is valid
		if !validators.IsValidUsername(username) {
			return c.JSON(fiber.Map{
				"availability": false,
			})
		}

		// Ensure there is no profanity
		detector := services.CustomProfanityDetector()
		if detector.IsProfane(username) {
			return c.JSON(fiber.Map{
				"availability": false,
			})
		}

		// Check if the username is in the reserved list
		if isUsernameReserved(username) {
			return c.JSON(fiber.Map{
				"availability": false,
			})
		}

		// Check if the username exists in the database
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check username availability.",
			})
		}

		// If the username exists, return false, else true
		isAvailable := !exists

		return c.JSON(fiber.Map{
			"availability": isAvailable,
		})
	}
}
