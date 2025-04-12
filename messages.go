package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type Message = models.Message

// GetChannelMessages retrieves up to 100 messages from a specific Discord channel.
//
// This function extracts the channel ID from the Fiber context and request parameters.
// It uses the DiscordGo session to retrieve the latest 100 messages from the specified channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the list of messages as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the messages cannot be retrieved.
//
// @Summary		Get Channel Messages
// @Description	Retrieve all messages from a specific channel.
// @Tags			Messages
// @Param			channelid	path		string	true	"Channel ID"
// @Success		200			{array}		Message
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/messages [get]
func GetChannelMessages(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	messages, err := s.ChannelMessages(channelID, 100, "", "", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.JSON(messages)
}

// GetChannelMessage retrieves a specific message from a Discord channel by its ID.
//
// This function extracts the channel ID and message ID from the Fiber context and request parameters.
// It uses the DiscordGo session to retrieve the specified message from the given channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the message as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the message cannot be retrieved.
//
// @Summary		Get Channel Message
// @Description	Retrieve a specific message by ID from a channel.
// @Tags			Messages
// @Param			channelid	path		string	true	"Channel ID"
// @Param			messageid	path		string	true	"Message ID"
// @Success		200			{object}	models.Message
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid} [get]
func GetChannelMessage(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")

	message, err := s.ChannelMessage(channelID, messageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve message: " + err.Error())
	}

	return c.JSON(message)
}

// SendChannelMessage sends a message to a specific Discord channel.
//
// This function extracts the channel ID from the Fiber context and request parameters.
// The message content is provided in the request body and parsed into a `discordgo.MessageSend` struct.
// It uses the DiscordGo session to send the message to the specified channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the sent message as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the message cannot be sent.
//
// @Summary		Send Channel Message
// @Description	Send a new message to a specific channel.
// @Tags			Messages
// @Param			channelid	path		string	true	"Channel ID"
// @Success		201			{object}	models.Message
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/messages [post]
func SendChannelMessage(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")

	var message discordgo.MessageSend
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	msg, err := s.ChannelMessageSendComplex(channelID, &message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send message: " + err.Error())
	}

	return c.JSON(msg)
}

// EditChannelMessage edits an existing message in a specific Discord channel.
//
// This function extracts the channel ID and message ID from the Fiber context and request parameters.
// The new message content is provided in the request body and parsed into a `discordgo.MessageEdit` struct.
// It uses the DiscordGo session to edit the message in the specified channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the edited message as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the message cannot be edited.
//
// @Summary		Edit Channel Message
// @Description	Edit a specific message in a channel by ID.
// @Tags			Messages
// @Param			channelid	path		string	true	"Channel ID"
// @Param			messageid	path		string	true	"Message ID"
// @Success		200			{object}	models.Message
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid} [patch]
func EditChannelMessage(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")

	var message discordgo.MessageEdit
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	message.ID = messageID
	message.Channel = channelID

	updatedMessage, err := s.ChannelMessageEditComplex(&message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to edit message: " + err.Error())
	}

	return c.JSON(updatedMessage)
}

// DeleteChannelMessage deletes a specific message from a Discord channel.
//
// This function extracts the channel ID and message ID from the Fiber context and request parameters.
// It uses the DiscordGo session to delete the specified message from the given channel.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content) if the message is successfully deleted.
//   - On failure, it returns an HTTP status 500 and an error message if the message cannot be deleted.
//
// @Summary		Delete Channel Message
// @Description	Delete a specific message in a channel by ID.
// @Tags			Messages
// @Param			channelid	path	string	true	"Channel ID"
// @Param			messageid	path	string	true	"Message ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid} [delete]
func DeleteChannelMessage(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")

	err := s.ChannelMessageDelete(channelID, messageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete message: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
