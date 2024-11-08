package disgm

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TokenMiddleware(disgm *Disgm, c *fiber.Ctx) error {
	//c.Locals("ID", "561234976788447232")
	//return c.Next()

	token := c.Get("Authorization")
	splToken := strings.Split(token, " ")
	if splToken[0] == "Bearer" {

		if disgm.opt.TokenStore != nil {
			tokens, err := disgm.opt.TokenStore.Load()

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
			}

			for k, v := range tokens {
				if v == splToken[1] {
					c.Locals("ID", k)
					return c.Next()
				}
			}
		}
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}
