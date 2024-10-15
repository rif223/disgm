package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type User = models.User

// GetBotUser is a handler function that retrieves the bot's user information
// from the Discord API and returns it as a JSON response.
//
// It uses the DiscordGo session (`s`) to fetch the bot's user object, which represents
// the bot account itself, identified by "@me".
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session that provides the connection to the Discord API.
//
// Returns:
//   - If successful, it returns the user object in JSON format with HTTP status 200.
//   - If an error occurs while retrieving the bot user, it returns an HTTP status 500
//     with an error message.
//	@Summary		Get Bot User
//	@Description	Retrieve the bot's user information.
//	@Tags			User
//	@Success		200	{object}	User
//	@Failure		500	{object}	error
//	@Router			/api/user [get]
func GetBotUser(c *fiber.Ctx, s *discordgo.Session) error {

	// Retrieve the bot user from the Discord API using the "@me" identifier
	user, err := s.User("@me")
	if err != nil {
		// Return a 500 status with an error message if the user retrieval fails
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve bot user: " + err.Error())
	}

	// Respond with the bot user information in JSON format
	return c.JSON(user)
}
