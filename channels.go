package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type ChannelArray = []models.Channel
type InviteArray = []models.Invite

// GetGuildChannels retrieves all channels for a specific guild.
//
// This function fetches the list of all channels in a guild using the guild ID stored in the
// Fiber context. The channels include text, voice, and category channels.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Returns:
//   - On success, it returns a JSON list of guild channels.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Guild Channels
// @Description	Retrieve all channels from the guild.
// @Tags			Channels
// @Success		200	{array}		ChannelArray
// @Failure		500	{object}	error
// @Router			/api/guild/channels [get]
func GetGuildChannels(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild channels: " + err.Error())
	}

	return c.JSON(channels)
}

// GetGuildChannel retrieves a specific channel by its ID.
//
// This function fetches the details of a specific guild channel using the channel ID passed in
// the request parameters.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel to retrieve.
//
// Returns:
//   - On success, it returns the channel details as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Guild Channel
// @Description	Retrieve a specific channel from the guild by ID.
// @Tags			Channels
// @Param			channelid	path		string	true	"Channel ID"
// @Success		200			{object}	models.Channel
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid} [get]
func GetGuildChannel(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	channel, err := s.Channel(channelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve channel: " + err.Error())
	}

	return c.JSON(channel)
}

// CreateGuildChannel creates a new channel in the specified guild.
//
// This function creates a new text, voice, or category channel in a guild based on the
// channel creation data provided in the request body. The guild ID is obtained from the
// Fiber context.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Context:
//   - ID: The guild ID is stored in the Fiber context under the key "ID".
//
// Request Body:
//   - The request body should contain the channel creation data in JSON format.
//
// Returns:
//   - On success, it returns the newly created channel as JSON.
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or an HTTP status 500 (Internal Server Error) if channel creation fails.
//
// @Summary		Create Guild Channel
// @Description	Create a new channel in the guild.
// @Tags			Channels
// @Success		201	{object}	models.Channel
// @Failure		500	{object}	error
// @Router			/api/guild/channels [post]
func CreateGuildChannel(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	var channelData discordgo.GuildChannelCreateData
	if err := c.BodyParser(&channelData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	channel, err := s.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create channel: " + err.Error())
	}

	return c.JSON(channel)
}

// UpdateGuildChannel updates an existing guild channel.
//
// This function updates the settings of a specific channel using the channel ID passed in the
// request parameters and the new channel settings provided in the request body.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel to update.
//
// Request Body:
//   - The request body should contain the new channel settings in JSON format.
//
// Returns:
//   - On success, it returns the updated channel as JSON.
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or an HTTP status 500 (Internal Server Error) if the update fails.
//
// @Summary		Update Guild Channel
// @Description	Update a specific channel in the guild.
// @Tags			Channels
// @Param			channelid	path		string	true	"Channel ID"
// @Success		200			{object}	models.Channel
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid} [patch]
func UpdateGuildChannel(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	var options *discordgo.ChannelEdit
	if err := c.BodyParser(&options); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	channel, err := s.ChannelEdit(channelID, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update channel positions: " + err.Error())
	}

	return c.JSON(channel)
}

// DeleteGuildChannel deletes a channel from a guild.
//
// This function deletes a channel using the channel ID passed in the request parameters.
// Deleting a channel will also remove all associated messages within the channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel to delete.
//
// Returns:
//   - On success, it returns the details of the deleted channel as JSON.
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Delete Guild Channel
// @Description	Delete a specific channel in the guild.
// @Tags			Channels
// @Param			channelid	path	string	true	"Channel ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid} [delete]
func DeleteGuildChannel(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	channel, err := s.ChannelDelete(channelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete channel: " + err.Error())
	}

	return c.JSON(channel)
}

// GetGuildChannelInvites retrieves all active invites for a specific channel.
//
// This function fetches all invite links associated with a given channel within a guild.
// The channel ID is passed as a path parameter in the request URL.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel whose invites should be retrieved.
//
// Returns:
//   - On success, returns a JSON array of invite objects.
//   - On failure, returns HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Get Channel Invites
// @Description	Retrieve all invites for a specific channel in the guild.
// @Tags			Channels
// @Param			channelid	path		string	true	"Channel ID"
// @Success		200			{array}		InviteArray
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/invites [get]
func GetGuildChannelInvites(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	invites, err := s.ChannelInvites(channelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve channel invites: " + err.Error())
	}

	return c.JSON(invites)
}

// CreateGuildChannelInvite creates an invite link for a specific channel.
//
// This function creates an invite for a specified channel within a guild. The channel ID is obtained
// from the request parameters, and the invite settings (e.g., max age, max uses) are provided in the
// request body as JSON.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel to create an invite for.
//
// Request Body:
//   - A JSON object containing invite parameters such as max_age, max_uses, temporary, and unique.
//
// Returns:
//   - On success, returns the created invite as JSON.
//   - On failure, returns HTTP 400 if the body is invalid, or HTTP 500 if the invite creation fails.
//
// @Summary		Create Channel Invite
// @Description	Create an invite for a specific channel in the guild.
// @Tags			Channels
// @Param			channelid	path		string					true	"Channel ID"
// @Param			inviteData	body		models.Invite	true	"Invite creation data"
// @Success		201			{object}	models.Invite
// @Failure		400			{object}	error
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/invites [post]
func CreateGuildChannelInvite(c *fiber.Ctx, s *discordgo.Session) error {
	//guildID := c.Locals("ID").(string)
	channelID := c.Params("channelid")

	var inviteData discordgo.Invite
	if err := c.BodyParser(&inviteData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	invite, err := s.ChannelInviteCreate(channelID, inviteData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create invite: " + err.Error())
	}

	return c.JSON(invite)
}

// EditChannelPermissions updates the permission overwrites for a channel.
//
// This function modifies the permissions for a channel by applying permission overwrites
// for a specific user or role, using the channel ID and overwrite ID provided in the request parameters.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel where permissions will be edited.
//   - overwriteid: The ID of the user or role whose permissions will be overwritten.
//
// Request Body:
//   - The request body should contain the permission overwrite details in JSON format.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or an HTTP status 500 (Internal Server Error) if permission updates fail.
//
// @Summary		Edit Channel Permissions
// @Description	Edit permissions for a specific channel in the guild.
// @Tags			Channels
// @Param			channelid	path	string	true	"Channel ID"
// @Param			overwriteid	path	string	true	"Overwrite ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/permissions/{overwriteid} [put]
func EditChannelPermissions(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	overwriteID := c.Params("overwriteid")

	var perm discordgo.PermissionOverwrite
	if err := c.BodyParser(&perm); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	err := s.ChannelPermissionSet(channelID, overwriteID, perm.Type, perm.Allow, perm.Deny)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to edit channel permissions: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteChannelPermissions removes a permission overwrite for a channel.
//
// This function deletes an existing permission overwrite for a channel using the channel ID
// and overwrite ID provided in the request parameters.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - channelid: The ID of the channel where permissions will be deleted.
//   - overwriteid: The ID of the user or role whose permission overwrite will be deleted.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 (Internal Server Error) with an error message.
//
// @Summary		Delete Channel Permissions
// @Description	Delete a specific permission overwrite for a channel.
// @Tags			Channels
// @Param			channelid	path	string	true	"Channel ID"
// @Param			overwriteid	path	string	true	"Overwrite ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/permissions/{overwriteid} [delete]
func DeleteChannelPermissions(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	overwriteID := c.Params("overwriteid")

	err := s.ChannelPermissionDelete(channelID, overwriteID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete channel permissions: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
