package disgm

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/rif223/disgm/store"

	_ "github.com/rif223/disgm/docs"
)

// Options contains the configuration for the disgm package.
type Options struct {
	DisableStartupMessage bool
	DisableLogger         bool
	TokenStore            store.TokenStore // A map of valid tokens for authentication.
}

// defaultOptions defines the default configuration for the disgm package.
var defaultOptions = Options{
	DisableStartupMessage: false,
	DisableLogger:         false,
}

// Disgm is the main structure for the package, containing the Discord session and the Fiber server.
type Disgm struct {
	opt   *Options           // Options for the application.
	s     *discordgo.Session // The DiscordGo session for interacting with the Discord API.
	fiber *fiber.App         // The Fiber application for the web server.
}

// New creates a new instance of Disgm with the specified DiscordGo session and options.
//
// Parameters:
//   - s: *discordgo.Session – The DiscordGo session used for interacting with the Discord API.
//   - options: ...Options – Optional configuration settings for the server.
//
// Returns:
//   - *Disgm: A new instance of Disgm.
//   - error: An error that may have occurred during initialization.

// @title			Discord Guild Management API
// @version		1.0
// @description	API for managing Discord guilds using DiscordGo and Fiber.
// @host			localhost:90
func New(s *discordgo.Session, options ...Options) (d *Disgm, err error) {

	opt := &defaultOptions

	if len(options) > 0 {
		o := options[0] // Gets the custom options.

		if o.TokenStore != nil {
			opt.TokenStore = o.TokenStore // Sets the valid tokens if specified.
		}
		if o.DisableStartupMessage {
			opt.DisableStartupMessage = o.DisableStartupMessage
		}
		if o.DisableLogger {
			opt.DisableLogger = o.DisableLogger
		}
	}

	app := fiber.New(fiber.Config{
		AppName:               "Disgm",
		DisableStartupMessage: opt.DisableStartupMessage,
		ProxyHeader:           "X-Forwarded-For", // Sets the proxy header for IP forwarding.
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Printf("Error: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error") // Returns an error status.
		},
	})

	d = &Disgm{
		opt:   opt, // Sets the default options.
		s:     s,   // Sets the DiscordGo session.
		fiber: app, // Sets the Fiber application.
	}

	// Configures CORS and logger middleware.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	if !opt.DisableLogger {
		app.Use(logger.New()) // Adds the logger.
	}

	// Middleware for token validation.
	app.Use(func(c *fiber.Ctx) error {
		return TokenMiddleware(d, c)
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	return
}

// Register Api Router
func (d *Disgm) RegisterApiRouter() {
	d.fiber.Route("/api", func(r fiber.Router) {
		Router(r, d.s) // Registers the API routes.
	})
}

// @Summary		Register WebSocket
// @Description	Sets up the WebSocket connection to handle Discord events and messages.
// @Tags			WebSocket
// @Produce		json
// @Router			/ws [get]
func (d *Disgm) RegisterWebSocket() {
	registerDiscordHandlers(d.s) // Registers the Discord handlers for events.

	// Sets the WebSocket connection.
	d.fiber.Get("/ws", websocket.New(func(c *websocket.Conn) {
		ID := c.Locals("ID").(string) // Retrieves the ID from the local context.
		WebSocket(c, ID)              // Handles the WebSocket connection.
	}))
}

// registerDiscordHandlers registers handlers for Discord events.
//
// This function adds an event handler that responds to various Discord events
// and processes the corresponding data.
//
// Parameters:
//   - s: *discordgo.Session – The DiscordGo session for interacting with the Discord API.
func registerDiscordHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, e *discordgo.Event) {
		// List of relevant events to handle.
		events := []string{
			"GUILD_UPDATE",
			"VOICE_STATE_UPDATE",
			"GUILD_MEMBER_ADD",
			"GUILD_MEMBER_UPDATE",
			"GUILD_MEMBER_REMOVE",
			"GUILD_BAN_ADD",
			"GUILD_BAN_REMOVE",
			"CHANNEL_CREATE",
			"CHANNEL_UPDATE",
			"CHANNEL_DELETE",
			"GUILD_ROLE_CREATE",
			"GUILD_ROLE_UPDATE",
			"GUILD_ROLE_DELETE",
			"MESSAGE_CREATE",
			"MESSAGE_UPDATE",
			"MESSAGE_DELETE",
			"MESSAGE_REACTION_ADD",
			"MESSAGE_REACTION_REMOVE",
			"MESSAGE_REACTION_REMOVE_ALL",
			"INTERACTION_CREATE",
		}

		// Checks if the event is in the list of processed events.
		if slices.Contains(events, e.Type) {
			var data map[string]interface{}

			err := json.Unmarshal(e.RawData, &data) // Converts the raw event data into a map.
			if err != nil {
				log.Printf("error: %v", err) // Logs errors when processing event data.
				return
			}

			if guildID, ok := data["guild_id"].(string); ok {
				EventCall(guildID, e.Type, data) // Calls the EventCall function with the relevant data.
			} else {
				fmt.Println("guild_id not found") // Logs if guild_id is not found.
			}
		}
	})
}

// Listen starts the Fiber server on the specified port.
//
// This method belongs to the `Disgm` type and initializes an HTTP server using the Fiber framework.
// By default, the port is set to ":90" if no other port is specified.
//
// Parameters:
//   - port (string): The port on which the server should listen. Defaults to ":90" if left empty.
//
// Return:
//   - error: Returns an error if the server fails to start.
//
// Functionality:
//   - Starts the server in a separate goroutine using Fiber (`app.Listen(port)`) to avoid blocking execution
//     and logs any errors encountered during startup.
//   - On success, logs a message indicating the actual port the server is listening on.
func (d *Disgm) Listen(port ...string) (err error) {
	if len(port) == 0 || port[0] == "" {
		port = append(port, ":90")
	}

	// Starts the Fiber server in a separate goroutine
	go func() {
		if err = d.fiber.Listen(port[0]); err != nil {
			log.Printf("Failed to start Fiber server: %v", err) // Logs any startup errors
		}
	}()
	log.Printf("Server started at port: %v", strings.Split(port[0], ":")[1]) // Logs startup message
	return err
}
