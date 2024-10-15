package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/rif223/disgm/models"
)

type Role = models.Role

// GetGuildRoles retrieves all roles of a specific guild.
//
// This function extracts the guild ID from the Fiber context's locals and uses
// the DiscordGo session to fetch all roles in the guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the roles as a JSON array with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the roles cannot be retrieved.
// @Summary		Get all roles in a guild
// @Description	Retrieve all roles of a specific guild using the guild ID.
// @Tags			Roles
// @Success		200	{array}		Role
// @Failure		500	{object}	error
// @Router			/api/guild/roles [get]
func GetGuildRoles(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild roles: " + err.Error())
	}

	return c.JSON(roles)
}

// GetGuildRole retrieves a specific role from a guild.
//
// This function extracts the guild ID and role ID from the Fiber context and request
// parameters, respectively, and uses the DiscordGo session to fetch the role.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the role object as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be retrieved.
// @Summary		Get a specific role in a guild
// @Description	Retrieve a specific role from a guild by its role ID.
// @Tags			Roles
// @Param			roleid	path		string	true	"ID of the role to retrieve"
// @Success		200		{object}	models.Role
// @Failure		500		{object}	error
// @Router			/api/guild/roles/{roleid} [get]
func GetGuildRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	roleID := c.Params("roleid")

	role, err := s.State.Role(guildID, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve role: " + err.Error())
	}

	return c.JSON(role)
}

// CreateGuildRole creates a new role in a guild.
//
// This function extracts the guild ID from the Fiber context and parses the request
// body to get role parameters. It uses the DiscordGo session to create a new role in the guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the created role as JSON with HTTP status 201.
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be created.
// @Summary		Create a new role in a guild
// @Description	Create a new role in a guild using the provided role parameters.
// @Tags			Roles
// @Param			body	body		models.RoleParams	true	"Role parameters"
// @Success		201		{object}	models.Role
// @Failure		500		{object}	error
// @Router			/api/guild/roles [post]
func CreateGuildRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	var roleData discordgo.RoleParams
	if err := c.BodyParser(&roleData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	role, err := s.GuildRoleCreate(guildID, &roleData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create role: " + err.Error())
	}

	return c.JSON(role)
}

// UpdateGuildRolePositions reorders the roles in a guild.
//
// This function extracts the guild ID from the Fiber context and parses the request body
// for the new role positions. It uses the DiscordGo session to update the role positions in the guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the updated roles as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the role positions cannot be updated.
// @Summary		Update role positions in a guild
// @Description	Reorder the roles in a guild based on the provided positions.
// @Tags			Roles
// @Param			body	body		[]models.Role	true	"New role positions"
// @Success		200		{array}		models.Role
// @Failure		500		{object}	error
// @Router			/api/guild/roles [patch]
func UpdateGuildRolePositions(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)

	var positions []*discordgo.Role
	if err := c.BodyParser(&positions); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	roles, err := s.GuildRoleReorder(guildID, positions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update role positions: " + err.Error())
	}

	return c.JSON(roles)
}

// UpdateGuildRole updates a specific role in a guild.
//
// This function extracts the guild ID and role ID from the Fiber context and request
// parameters, and parses the request body for role data. It uses the DiscordGo session to update the role.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns the updated role as JSON with HTTP status 200.
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be updated.
// @Summary		Update a specific role in a guild
// @Description	Update a specific role in a guild using the provided role data.
// @Tags			Roles
// @Param			roleid	path		string				true	"ID of the role to update"
// @Param			body	body		models.RoleParams	true	"Updated role parameters"
// @Success		200		{object}	models.Role
// @Failure		500		{object}	error
// @Router			/api/guild/roles/{roleid} [patch]
func UpdateGuildRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	roleID := c.Params("roleid")

	var roleData *discordgo.RoleParams
	if err := c.BodyParser(&roleData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	role, err := s.GuildRoleEdit(guildID, roleID, roleData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update role: " + err.Error())
	}

	return c.JSON(role)
}

// DeleteGuildRole deletes a specific role from a guild.
//
// This function extracts the guild ID and role ID from the Fiber context and request
// parameters, and uses the DiscordGo session to delete the role from the guild.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Returns:
//   - On success, it returns an HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 500 and an error message if the role cannot be deleted.
// @Summary		Delete a role from a guild
// @Description	Delete a specific role from a guild using its role ID.
// @Tags			Roles
// @Param			roleid	path	string	true	"ID of the role to delete"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/roles/{roleid} [delete]
func DeleteGuildRole(c *fiber.Ctx, s *discordgo.Session) error {
	guildID := c.Locals("ID").(string)
	roleID := c.Params("roleid")

	err := s.GuildRoleDelete(guildID, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete role: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
