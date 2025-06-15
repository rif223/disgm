package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type Guild = models.Guild

// GetGuild retrieves the details of a Discord guild.
//
// This function fetches the guild information using the guild ID, which is extracted from
// the request context. It uses the DiscordGo session to request guild details from the Discord API.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns the guild details as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Guild
// @Description	Retrieve the guild information.
// @Tags			Guild
// @Success		200	{object}	Guild
// @Failure		500	{object}	error
// @Router			/api/guild [get]
func GetGuild(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	guild, err := s.Guild(guildID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild: " + err.Error())
	}

	return c.JSON(guild)
}

// UpdateGuild updates the settings of a Discord guild.
//
// This function modifies the guild using the provided guild ID from the request context
// and the new guild data parsed from the request body. It utilizes the DiscordGo session
// to perform the update via the Discord API.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Request Body:
//   - Expects a JSON object that conforms to the discordgo.GuildParams struct,
//     containing the fields to be updated (e.g., name, region, etc.).
//
// Returns:
//   - On success, it returns the updated guild details as JSON.
//   - On failure:
//       - If the request body is invalid, it returns an HTTP status 400 (Bad Request).
//       - If the Discord API request fails, it returns an HTTP status 500 (Internal Server Error).
//
// @Summary		Update Guild
// @Description	Update the settings of a Discord guild.
// @Tags			Guild
// @Accept			json
// @Produce		json
// @Param			guild	body	models.GuildParams	true	"Guild parameters to update"
// @Success		200		{object}	Guild
// @Failure		400		{string}	string	"Invalid request body"
// @Failure		500		{string}	string	"Failed to update guild"
// @Router			/api/guild [patch]
func UpdateGuild(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	var guildData discordgo.GuildParams
	if err := c.BodyParser(&guildData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	guild, err := s.GuildEdit(guildID, &guildData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update guild: " + err.Error())
	}

	return c.JSON(guild)
}

// GetGuildBans retrieves the list of bans for a Discord guild.
//
// This function fetches a list of banned members from a guild by using the guild ID,
// which is stored in the request context. It returns up to 100 bans at a time.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns the list of bans as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Guild Bans
// @Description	Retrieve all banned users from the guild.
// @Tags			Bans
// @Success		200	{array}		models.GuildBan
// @Failure		500	{object}	error
// @Router			/api/guild/bans [get]
func GetGuildBans(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	bans, err := s.GuildBans(guildID, 100, "", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild bans: " + err.Error())
	}

	return c.JSON(bans)
}

// GetGuildBan retrieves information about a specific banned member in a guild.
//
// This function fetches the details of a specific ban using the guild ID and user ID, both of which
// are retrieved from the request parameters and context.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - userid: The ID of the banned user.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns the details of the banned member as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Guild Ban
// @Description	Retrieve a specific banned user by user ID.
// @Tags			Bans
// @Param			userid	path		string	true	"User ID"
// @Success		200		{object}	models.GuildBan
// @Failure		500		{object}	error
// @Router			/api/guild/bans/{userid} [get]
func GetGuildBan(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	userID := c.Params("userid")

	ban, err := s.GuildBan(guildID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild ban: " + err.Error())
	}

	return c.JSON(ban)
}

// AddGuildBan adds a ban to a user in a Discord guild.
//
// This function bans a member from a guild by using the guild ID and user ID.
// It also allows for specifying a reason and the number of days of message history to delete.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - userid: The ID of the user to be banned.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Request Body:
//   - The request body should contain a JSON object with the fields "reason" (string) and
//     "delete_message_days" (int).
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or HTTP status 500 (Internal Server Error) if the ban creation fails.
//
// @Summary		Add Guild Ban
// @Description	Ban a user from the guild.
// @Tags			Bans
// @Param			userid	path	string	true	"User ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/bans/{userid} [put]
func AddGuildBan(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	userID := c.Params("userid")

	var banData struct {
		Reason            string `json:"reason"`
		DeleteMessageDays int    `json:"delete_message_days"`
	}
	if err := c.BodyParser(&banData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	err := s.GuildBanCreateWithReason(guildID, userID, banData.Reason, banData.DeleteMessageDays)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add guild ban: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// RemoveGuildBan removes a ban from a user in a Discord guild.
//
// This function lifts the ban on a specific member by using the guild ID and user ID
// from the request parameters and context.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - userid: The ID of the user whose ban should be removed.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Remove Guild Ban
// @Description	Remove a ban for a user in the guild.
// @Tags			Bans
// @Param			userid	path	string	true	"User ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/bans/{userid} [delete]
func RemoveGuildBan(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	userID := c.Params("userid")

	err := s.GuildBanDelete(guildID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to remove guild ban: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// BulkBanMembers bans multiple members from a Discord guild.
//
// This function bans multiple users at once in a guild. The user IDs are provided in the request body
// as an array, and each user is banned using the DiscordGo session.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The ID of the guild is stored in the Fiber context under the key "ID".
//
// Request Body:
//   - The body should contain an array of user IDs (strings) to be banned.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or HTTP status 500 (Internal Server Error) if banning any user fails.
//
// @Summary		Bulk Ban Members
// @Description	Ban multiple users in the guild at once.
// @Tags			Bans
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/bulk-ban [post]
func BulkBanMembers(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	var userIDs []string
	if err := c.BodyParser(&userIDs); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	for _, userID := range userIDs {
		err := s.GuildBanCreate(guildID, userID, 0)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to ban user: " + err.Error())
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}
