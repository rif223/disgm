package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type ApplicationCommandArray = []models.ApplicationCommand

// GetGuildApplicationCommands retrieves all application commands for a specific guild.
//
// This function fetches the list of application commands registered for a guild, using
// the guild ID from the Fiber context and the bot's application ID retrieved from the Discord API.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns a JSON list of application commands.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
// @Summary		Get Guild Application Commands
// @Description	Retrieve all guild application commands.
// @Tags			Commands
// @Success		200	{array}		ApplicationCommandArray
// @Failure		500	{object}	error
// @Router			/api/guild/commands [get]
func GetGuildApplicationCommands(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	user, _ := s.User("@me") // Retrieves the bot's application user

	cmd, err := s.ApplicationCommands(user.ID, guildID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve cmds: " + err.Error())
	}

	return c.JSON(cmd)
}

// GetGuildApplicationCommand retrieves a specific application command for a guild.
//
// This function fetches details for a specific application command by its ID, using the
// guild ID from the Fiber context and the bot's application ID from the Discord API.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - cmdid: The ID of the application command to retrieve.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns the application command details as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
// @Summary		Get Guild Application Command
// @Description	Retrieve a specific guild application command by ID.
// @Tags			Commands
// @Param			cmdid	path		string	true	"Command ID"
// @Success		200		{object}	models.ApplicationCommand
// @Failure		500		{object}	error
// @Router			/api/guild/commands/{cmdid} [get]
func GetGuildApplicationCommand(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	user, _ := s.User("@me") // Retrieves the bot's application user
	cmdID := c.Params("cmdid")

	cmd, err := s.ApplicationCommand(user.ID, guildID, cmdID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve cmd: " + err.Error())
	}

	return c.JSON(cmd)
}

// CreateGuildApplicationCommand registers a new application command for a specific guild.
//
// This function creates a new application command in the specified guild by providing the
// command data in the request body. The command is associated with the bot's application ID.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Request Body:
//   - The request body should contain the application command data in JSON format.
//
// Returns:
//   - On success, it returns the newly created application command as JSON.
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or an HTTP status 500 (Internal Server Error) if command creation fails.
// @Summary		Create Guild Application Command
// @Description	Create a new guild application command.
// @Tags			Commands
// @Success		201	{object}	models.ApplicationCommand
// @Failure		500	{object}	error
// @Router			/api/guild/commands [post]
func CreateGuildApplicationCommand(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	user, _ := s.User("@me") // Retrieves the bot's application user

	var ac *discordgo.ApplicationCommand
	if err := c.BodyParser(&ac); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	cmd, err := s.ApplicationCommandCreate(user.ID, guildID, ac)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create cmd: " + err.Error())
	}

	return c.JSON(cmd)
}

// DeleteGuildApplicationCommand deletes a specific application command from a guild.
//
// This function removes an existing application command from a guild, using the guild ID
// from the Fiber context and the command ID provided in the request parameters.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - cmdid: The ID of the application command to delete.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
// @Summary		Delete Guild Application Command
// @Description	Delete a guild application command by ID.
// @Tags			Commands
// @Param			cmdid	path	string	true	"Command ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/commands/{cmdid} [delete]
func DeleteGuildApplicationCommand(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	user, _ := s.User("@me") // Retrieves the bot's application user
	cmdID := c.Params("cmdid")

	err := s.ApplicationCommandDelete(user.ID, guildID, cmdID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete cmd: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
