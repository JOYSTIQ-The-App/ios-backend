package middleware

import "github.com/gofiber/fiber/v2"

func IPFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		allowedIPs := map[string]bool{
			"127.0.0.1": true, // localhost
			// Add other IPs as needed
		}

		if _, allowed := allowedIPs[ip]; !allowed {
			return c.Status(fiber.StatusForbidden).SendString("Not Allowed")
		}

		return c.Next()
	}
}
