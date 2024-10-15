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
	Name string      `json:"name"` // Name of the event (e.g., "MESSAGE_CREATE")
	Data interface{} `json:"data"` // Event data (can be of any type)
}

// A map to keep track of connected clients. The map key is the WebSocket connection,
// and the value is the client's unique ID.
var clients = make(map[*websocket.Conn]string)

// WebSocket function manages the lifecycle of a WebSocket connection.
// It registers the client, sends a welcome message, and listens for incoming messages.
func WebSocket(conn *websocket.Conn, id string) {
	defer func() {
		// Close the connection and clean up when the function exits
		conn.Close()
	}()

	// Register the client with their unique ID
	clients[conn] = id
	log.Printf("Client connected: %s", id)

	// Send a welcome message to the client
	conn.WriteMessage(websocket.TextMessage, []byte("Welcome! You are connected."))

	// Handle incoming messages from the client
	handleMessages(conn, id)
}

// handleMessages continuously listens for messages from the connected client
// and logs the received messages. It also handles client disconnections.
func handleMessages(conn *websocket.Conn, id string) {
	defer func() {
		// Close the connection and remove the client from the map on disconnect
		conn.Close()
		delete(clients, conn)
		log.Printf("Client disconnected: %s", id)
	}()

	// Loop to continuously read messages from the WebSocket connection
	for {
		_, msg, err := conn.ReadMessage() // Read the message from the client
		if err != nil {
			// Log any errors (like client disconnection or message read error)
			log.Printf("error: %v", err)
			break
		}
		// Log the message along with the client ID
		log.Printf("%s: %s", id, msg)
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
				return fmt.Errorf("error marshalling message: %v", err)
			}

			// Write the JSON-encoded event to the client's WebSocket connection
			return client.WriteMessage(websocket.TextMessage, eventBytes)
		}
	}
	// Return nil if no matching client is found
	return nil
}
