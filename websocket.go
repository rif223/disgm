package disgm

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
)

// Event struct defines the structure of an event that is sent to clients over WebSocket.
// It contains the event name and the associated data.
type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type WS struct {
	conn *websocket.Conn // The WebSocket connection for real-time communication.
	id   string          // The unique ID of the client connected via WebSocket.
}

// A map to keep track of connected clients. The map key is the WebSocket connection,
// and the value is the client's unique ID.
var clients = make(map[*websocket.Conn]string)

// WebSocket function manages the lifecycle of a WebSocket connection.
// It registers the client, sends a welcome message, and listens for incoming messages.
func NewWebSocket(conn *websocket.Conn, id string) (*WS, error) {
	defer func() {
		conn.Close()
	}()

	// Register the client with their unique ID
	clients[conn] = id
	// time id status ip method path msg
	log.Printf("| %s | %s | %s | %s | %s | %s\n",
		id,
		"\u001b[92m OK\u001b[0m",
		conn.IP(),
		"\u001b[94m WS\u001b[0m",
		"/ws",
		"Client connected!",
	)

	// Send a welcome message to the client
	err := conn.WriteMessage(websocket.TextMessage, []byte("You are connected."))
	if err != nil {
		return nil, err
	}

	return &WS{
		conn: conn,
		id:   id,
	}, nil
}

// handleMessages continuously listens for messages from the connected client
// and logs the received messages. It also handles client disconnections.
func (ws *WS) handleMessages(messageHandlerFunc func(ws *WS, id string, msg []byte)) {
	defer func() {
		// Close the connection and remove the client from the map on disconnect
		ws.conn.Close()
		delete(clients, ws.conn)
		log.Printf("| %s | %s | %s | %s | %s | %s\n",
			ws.id,
			"\u001b[92m OK\u001b[0m",
			ws.conn.IP(),
			"\u001b[94m WS\u001b[0m",
			"/ws",
			"Client disconnected!",
		)
	}()

	// Loop to continuously read messages from the WebSocket connection
	for {
		_, msg, err := ws.conn.ReadMessage() // Read the message from the client
		if err != nil {
			// Log any errors (like client disconnection or message read error)
			log.Printf("| %s | %s | %s | %s | %s | %s\n",
				ws.id,
				"\u001b[91mERR\u001b[0m",
				ws.conn.IP(),
				"\u001b[94m WS\u001b[0m",
				"/ws",
				err.Error(),
			)
			break
		}
		messageHandlerFunc(ws, ws.id, msg) // Call the message handler function with the client ID and message
	}
}

// EventCall is used to send an event to a specific client identified by the ID.
// It marshals the event data to JSON and sends it via WebSocket to the client.
func EventCall(id string, name string, data interface{}) error {
	// Iterate over all connected clients
	for client, gid := range clients {
		// Send the event to the client with the matching ID
		if gid == id {
			// Create an Event struct with the event name and data
			event := Event{
				Name: name,
				Data: data,
			}

			// Marshal the event into JSON format
			eventBytes, err := json.Marshal(event)
			if err != nil {
				// Return an error if JSON marshalling fails
				return fmt.Errorf("marshalling message: %v", err)
			}

			// Write the JSON-encoded event to the client's WebSocket connection
			return client.WriteMessage(websocket.TextMessage, eventBytes)
		}
	}
	return nil
}
