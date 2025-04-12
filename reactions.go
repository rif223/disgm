package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type UserArray = []models.User

// GetMessageReactions retrieves the users who reacted to a specific message with a given emoji.
//
// This function extracts the channel ID, message ID, and emoji ID from the Fiber context and request parameters.
// It uses the DiscordGo session to retrieve the list of users who reacted with the specified emoji.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the list of users who reacted with the emoji as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the reactions cannot be retrieved.
//
// @Summary		Get Message Reactions
// @Description	Retrieve all reactions from a specific message in a channel.
// @Tags			Reactions
// @Param			channelid	path		string	true	"Channel ID"
// @Param			messageid	path		string	true	"Message ID"
// @Param			emojiid		path		string	true	"Emoji ID"
// @Success		200			{array}		UserArray
// @Failure		500			{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid}/reactions/{emojiid} [get]
func GetMessageReactions(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")
	emojiID := c.Params("emojiid")

	users, err := s.MessageReactions(channelID, messageID, emojiID, 100, "", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.JSON(users)
}

// CreateMessageReaction adds a reaction to a specific message with a given emoji.
//
// This function extracts the channel ID, message ID, and emoji ID from the Fiber context and request parameters.
// It uses the DiscordGo session to add a reaction to the specified message with the given emoji.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content) if the reaction is successfully added.
//   - On failure, it returns an HTTP status 500 and an error message if the reaction cannot be added.
//
// @Summary		Create Message Reaction
// @Description	Add a reaction to a specific message in a channel.
// @Tags			Reactions
// @Param			channelid	path	string	true	"Channel ID"
// @Param			messageid	path	string	true	"Message ID"
// @Param			emojiid		path	string	true	"Emoji ID"
// @Success		201
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid}/reactions/{emojiid} [put]
func CreateMessageReaction(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")
	emojiID := c.Params("emojiid")

	err := s.MessageReactionAdd(channelID, messageID, emojiID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteMessageReaction removes a specific user's reaction from a message.
//
// This function extracts the channel ID, message ID, emoji ID, and user ID from the Fiber context and request parameters.
// It uses the DiscordGo session to remove a user's reaction to a message.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content) if the reaction is successfully removed.
//   - On failure, it returns an HTTP status 500 and an error message if the reaction cannot be removed.
//
// @Summary		Delete Message Reaction
// @Description	Delete a user's reaction from a specific message in a channel.
// @Tags			Reactions
// @Param			channelid	path	string	true	"Channel ID"
// @Param			messageid	path	string	true	"Message ID"
// @Param			emojiid		path	string	true	"Emoji ID"
// @Param			userid		path	string	true	"User ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid}/reactions/{emojiid}/{userid} [delete]
func DeleteMessageReaction(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")
	emojiID := c.Params("emojiid")
	userID := c.Params("userid")

	err := s.MessageReactionRemove(channelID, messageID, emojiID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteAllMessageReaction removes all reactions from a message.
//
// This function extracts the channel ID and message ID from the Fiber context and request parameters.
// It uses the DiscordGo session to remove all reactions from the specified message.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content) if all reactions are successfully removed.
//   - On failure, it returns an HTTP status 500 and an error message if the reactions cannot be removed.
//
// @Summary		Delete All Message Reactions
// @Description	Remove all reactions from a specific message in a channel.
// @Tags			Reactions
// @Param			channelid	path	string	true	"Channel ID"
// @Param			messageid	path	string	true	"Message ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid}/reactions [delete]
func DeleteAllMessageReaction(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")

	err := s.MessageReactionsRemoveAll(channelID, messageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteMessageReactionEmoji removes all reactions for a specific emoji from a message.
//
// This function extracts the channel ID, message ID, and emoji ID from the Fiber context and request parameters.
// It uses the DiscordGo session to remove all reactions for the specified emoji from the message.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content) if all reactions for the emoji are successfully removed.
//   - On failure, it returns an HTTP status 500 and an error message if the reactions cannot be removed.
//
// @Summary		Delete Message Reaction Emoji
// @Description	Remove a specific emoji reaction from a message in a channel.
// @Tags			Reactions
// @Param			channelid	path	string	true	"Channel ID"
// @Param			messageid	path	string	true	"Message ID"
// @Param			emojiid		path	string	true	"Emoji ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/channels/{channelid}/messages/{messageid}/reactions/{emojiid} [delete]
func DeleteMessageReactionEmoji(c *fiber.Ctx, s *discordgo.Session) error {
	channelID := c.Params("channelid")
	messageID := c.Params("messageid")
	emojiID := c.Params("emojiid")

	err := s.MessageReactionsRemoveEmoji(channelID, messageID, emojiID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve messages: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
