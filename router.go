package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router, s *discordgo.Session) {

	router.Get("/user", func(c *fiber.Ctx) error {
		return GetBotUser(c, s)
	})

	router.Get("/guild", func(c *fiber.Ctx) error {
		return GetGuild(c, s)
	})

	router.Post("/guild/interactions/:interactionid/:interactiontoken/callback", func(c *fiber.Ctx) error {
		return CreateInteractionCallback(c, s)
	})

	router.Get("/guild/commands", func(c *fiber.Ctx) error {
		return GetGuildApplicationCommands(c, s)
	})

	router.Get("/guild/commands/:cmdid", func(c *fiber.Ctx) error {
		return GetGuildApplicationCommand(c, s)
	})

	router.Post("/guild/commands", func(c *fiber.Ctx) error {
		return CreateGuildApplicationCommand(c, s)
	})

	router.Delete("/guild/commands/:cmdid", func(c *fiber.Ctx) error {
		return DeleteGuildApplicationCommand(c, s)
	})

	router.Get("/guild/bans", func(c *fiber.Ctx) error {
		return GetGuildBans(c, s)
	})

	router.Get("/guild/bans/:userid", func(c *fiber.Ctx) error {
		return GetGuildBan(c, s)
	})

	router.Put("/guild/bans/:userid", func(c *fiber.Ctx) error {
		return AddGuildBan(c, s)
	})

	router.Delete("/guild/bans/:userid", func(c *fiber.Ctx) error {
		return RemoveGuildBan(c, s)
	})

	router.Post("/guild/bulk-ban", func(c *fiber.Ctx) error {
		return BulkBanMembers(c, s)
	})

	router.Get("/guild/channels", func(c *fiber.Ctx) error {
		return GetGuildChannels(c, s)
	})

	router.Get("/guild/channels/:channelid", func(c *fiber.Ctx) error {
		return GetGuildChannel(c, s)
	})

	router.Post("/guild/channels", func(c *fiber.Ctx) error {
		return CreateGuildChannel(c, s)
	})

	router.Patch("/guild/channels/:channelid", func(c *fiber.Ctx) error {
		return UpdateGuildChannel(c, s)
	})

	router.Delete("/guild/channels/:channelid", func(c *fiber.Ctx) error {
		return DeleteGuildChannel(c, s)
	})

	router.Put("/guild/channels/:channelid/permissions/:overwriteid", func(c *fiber.Ctx) error {
		return EditChannelPermissions(c, s)
	})

	router.Delete("/guild/channels/:channelid/permissions/:overwriteid", func(c *fiber.Ctx) error {
		return DeleteChannelPermissions(c, s)
	})

	router.Get("/guild/channels/:channelid/messages", func(c *fiber.Ctx) error {
		return GetChannelMessages(c, s)
	})

	router.Get("/guild/channels/:channelid/messages/:messageid", func(c *fiber.Ctx) error {
		return GetChannelMessage(c, s)
	})

	router.Post("/guild/channels/:channelid/messages", func(c *fiber.Ctx) error {
		return SendChannelMessage(c, s)
	})

	router.Patch("/guild/channels/:channelid/messages/:messageid", func(c *fiber.Ctx) error {
		return EditChannelMessage(c, s)
	})

	router.Delete("/guild/channels/:channelid/messages/:messageid", func(c *fiber.Ctx) error {
		return DeleteChannelMessage(c, s)
	})

	router.Get("/guild/channels/:channelid/messages/:messageid/reactions/:emojiid", func(c *fiber.Ctx) error {
		return GetMessageReactions(c, s)
	})

	router.Put("/guild/channels/:channelid/messages/:messageid/reactions/:emojiid", func(c *fiber.Ctx) error {
		return CreateMessageReaction(c, s)
	})

	router.Delete("/guild/channels/:channelid/messages/:messageid/reactions/:emojiid/:userid", func(c *fiber.Ctx) error {
		return DeleteMessageReaction(c, s)
	})

	router.Get("/guild/channels/:channelid/messages/:messageid/reactions", func(c *fiber.Ctx) error {
		return DeleteAllMessageReaction(c, s)
	})

	router.Get("/guild/channels/:channelid/messages/:messageid/reactions/:emojiid", func(c *fiber.Ctx) error {
		return DeleteMessageReactionEmoji(c, s)
	})

	router.Get("/guild/members", func(c *fiber.Ctx) error {
		return GetGuildMembers(c, s)
	})

	router.Get("/guild/members/:memberid", func(c *fiber.Ctx) error {
		return GetGuildMember(c, s)
	})

	router.Patch("/guild/members/:memberid", func(c *fiber.Ctx) error {
		return UpdateGuildMember(c, s)
	})

	router.Delete("/guild/members/:memberid", func(c *fiber.Ctx) error {
		return KickMember(c, s)
	})

	router.Get("/guild/members/:memberid/roles", func(c *fiber.Ctx) error {
		return GetMemberRoles(c, s)
	})

	router.Put("/guild/members/:memberid/roles/:roleid", func(c *fiber.Ctx) error {
		return AddMemberRole(c, s)
	})

	router.Delete("/guild/members/:memberid/roles/:roleid", func(c *fiber.Ctx) error {
		return RemoveMemberRole(c, s)
	})

	router.Get("/guild/roles", func(c *fiber.Ctx) error {
		return GetGuildRoles(c, s)
	})

	router.Patch("/guild/roles", func(c *fiber.Ctx) error {
		return UpdateGuildRolePositions(c, s)
	})

	router.Get("/guild/roles/:roleid", func(c *fiber.Ctx) error {
		return GetGuildRole(c, s)
	})

	router.Post("/guild/roles/:roleid", func(c *fiber.Ctx) error {
		return CreateGuildRole(c, s)
	})

	router.Patch("/guild/roles/:roleid", func(c *fiber.Ctx) error {
		return UpdateGuildRole(c, s)
	})

	router.Delete("/guild/roles/:roleid", func(c *fiber.Ctx) error {
		return DeleteGuildRole(c, s)
	})
}
