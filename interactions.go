package disgm

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

// CreateInteractionCallback handles the creation of a response to a Discord interaction.
//
// This function receives an interaction ID and interaction token from the request parameters,
// along with the interaction response data from the request body. It then sends the interaction
// response using the provided DiscordGo session.
//
// Parameters:
//   - c: *fiber.Ctx – The Fiber context used to handle HTTP requests and responses.
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//
// Request Parameters:
//   - interactionid: The ID of the interaction.
//   - interactiontoken: The token of the interaction.
//
// Request Body:
//   - The body should contain a valid `discordgo.InteractionResponse` object in JSON format.
//
// Returns:
//   - On success, it returns HTTP status 204 (No Content).
//   - On failure, it returns an HTTP status 400 (Bad Request) if the request body is invalid,
//     or HTTP status 500 (Internal Server Error) if there is a problem sending the response.
//
// @Summary		Create Interaction Callback
// @Description	Handle interaction callback for a specific interaction.
// @Tags			Interactions
// @Param			interactionid		path	string	true	"Interaction ID"
// @Param			interactiontoken	path	string	true	"Interaction Token"
// @Success		204
// @Failure		500	{object}	error
// @Router			/api/guild/interactions/{interactionid}/{interactiontoken}/callback [post]
func CreateInteractionCallback(c *fiber.Ctx, s *discordgo.Session) error {
	interactionID := c.Params("interactionid")
	interactionToken := c.Params("interactiontoken")

	var resp *discordgo.InteractionResponse
	if err := c.BodyParser(&resp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	err := NewInteractionRespond(s, interactionID, interactionToken, resp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve guild channels: " + err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// NewInteractionRespond sends a response to a Discord interaction.
//
// This function sends an interaction response to the Discord API using the interaction ID and token.
// If the response contains files (attachments), it uses multipart form data to handle them.
//
// Parameters:
//   - s: *discordgo.Session – The DiscordGo session used to interact with the Discord API.
//   - id: string – The interaction ID.
//   - token: string – The interaction token.
//   - resp: *discordgo.InteractionResponse – The response data to be sent for the interaction.
//   - options: (optional) Additional request options (e.g., custom headers).
//
// Returns:
//   - On success, it returns `nil`.
//   - On failure, it returns an error if there is an issue preparing or sending the request.
func NewInteractionRespond(s *discordgo.Session, id string, token string, resp *discordgo.InteractionResponse, options ...discordgo.RequestOption) error {
	endpoint := discordgo.EndpointInteractionResponse(id, token)

	if resp.Data != nil && len(resp.Data.Files) > 0 {
		_, body, err := discordgo.MultipartBodyWithJSON(resp, resp.Data.Files)
		if err != nil {
			return err
		}

		_, err = s.Request("POST", endpoint, body, options...)
		return err
	}

	_, err := s.RequestWithBucketID("POST", endpoint, *resp, endpoint, options...)
	return err
}
