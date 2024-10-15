package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type Member = models.Member

// GetGuildMembers retrieves a list of up to 1000 members from a specific Discord guild.
//
// This function extracts the guild ID from the Fiber context and uses the DiscordGo session to
// retrieve the guild members. It fetches up to 1000 members from the specified guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the list of guild members as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the members cannot be retrieved.
// @Summary		Get Guild Members
// @Description	Retrieve all members of the guild.
// @Tags			Members
// @Success		200	{array}		Member
// @Failure		500	{object}	error
// @Router			/api/guild/members [get]
func GetGuildMembers(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild members: " + err.Error())
	}

	return c.JSON(members)
}

// GetGuildMember retrieves a specific member from a Discord guild using their member ID.
//
// This function extracts the guild ID from the Fiber context and the member ID from the request parameters.
// It uses the DiscordGo session to retrieve the member from the specified guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the guild member as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the member cannot be retrieved.
// @Summary		Get Guild Member
// @Description	Retrieve a specific member from the guild by ID.
// @Tags			Members
// @Param			memberid	path		string	true	"Member ID"
// @Success		200			{object}	models.Member
// @Failure		500			{object}	error
// @Router			/api/guild/members/{memberid} [get]
func GetGuildMember(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")

	member, err := s.GuildMember(guildID, memberID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild member: " + err.Error())
	}

	return c.JSON(member)
}

// UpdateGuildMember modifies the settings of a guild member.
//
// This function extracts the guild ID and member ID from the Fiber context and request parameters.
// It parses the request body into a `discordgo.GuildMemberParams` struct and uses it to update
// the member's settings (e.g., nickname, roles, mute, etc.).
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the updated guild member as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the member cannot be updated.
// @Summary		Update Guild Member
// @Description	Update a specific member in the guild.
// @Tags			Members
// @Param			memberid	path		string	true	"Member ID"
// @Success		200			{object}	models.Member
// @Failure		500			{object}	error
// @Router			/api/guild/members/{memberid} [patch]
func UpdateGuildMember(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")

	var memberEdit discordgo.GuildMemberParams
	if err := c.BodyParser(&memberEdit); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	member, err := s.GuildMemberEdit(guildID, memberID, &memberEdit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update guild member: " + err.Error())
	}

	return c.JSON(member)
}

// GetMemberRoles retrieves the roles of a specific guild member.
//
// This function extracts the guild ID and member ID from the Fiber context and request parameters.
// It uses the DiscordGo session to retrieve the member's roles in the specified guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the list of roles assigned to the member as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the member roles cannot be retrieved.
// @Summary		Get Member Roles
// @Description	Retrieve all roles assigned to a specific member in the guild.
// @Tags			Roles
// @Param			memberid	path		string	true	"Member ID"
// @Success		200			{array}		models.Role
// @Failure		500			{object}	error
// @Router			/api/guild/members/{memberid}/roles [get]
func GetMemberRoles(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")

	member, err := s.GuildMember(guildID, memberID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve member roles: " + err.Error())
	}

	return c.JSON(member.Roles)
}

// AddMemberRole adds a role to a guild member.
//
// This function extracts the guild ID, member ID, and role ID from the Fiber context and request parameters.
// It uses the DiscordGo session to add the specified role to the member.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be added.
// @Summary		Add Member Role
// @Description	Add a role to a specific member in the guild.
// @Tags			Roles
// @Param			memberid	path	string	true	"Member ID"
// @Param			roleid		path	string	true	"Role ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/members/{memberid}/roles/{roleid} [put]
func AddMemberRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")
	roleID := c.Params("roleid")

	err := s.GuildMemberRoleAdd(guildID, memberID, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add role to member: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// RemoveMemberRole removes a role from a guild member.
//
// This function extracts the guild ID, member ID, and role ID from the Fiber context and request parameters.
// It uses the DiscordGo session to remove the specified role from the member.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be removed.
// @Summary		Remove Member Role
// @Description	Remove a role from a specific member in the guild.
// @Tags			Roles
// @Param			memberid	path	string	true	"Member ID"
// @Param			roleid		path	string	true	"Role ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/members/{memberid}/roles/{roleid} [delete]
func RemoveMemberRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")
	roleID := c.Params("roleid")

	err := s.GuildMemberRoleRemove(guildID, memberID, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to remove role from member: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// KickMember removes a member from the guild.
//
// This function extracts the guild ID and member ID from the Fiber context and request parameters.
// It uses the DiscordGo session to remove the specified member from the guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 and an error message if the member cannot be removed.
// @Summary		Kick Member
// @Description	Remove a member from the specified guild.
// @Tags			Members
// @Param			memberid	path	string	true	"Member ID"
// @Success		204
// @Failure		500	{object}	error
// @Router			/guilds/{guildid}/members/{memberid} [delete]
func KickMember(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	memberID := c.Params("memberid")

	err := s.GuildMemberDelete(guildID, memberID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to kick member: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
