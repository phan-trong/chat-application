package websocket

import (
	"fmt"

	"github.com/google/uuid"
)

const welcomeMessage = "%s joined the room"

type RoomChat struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *MessageChat
	Private    bool `json:"private"`
}

// NewRoom creates a new Room
func NewRoom(name string, private bool) *RoomChat {
	return &RoomChat{
		ID:         uuid.New(),
		Name:       name,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *MessageChat),
		Private:    private,
	}
}

// RunRoom runs our room, accepting various requests
func (room *RoomChat) RunRoom() {
	for {
		select {

		case client := <-room.register:
			room.registerClientInRoom(client)

		case client := <-room.unregister:
			room.unregisterClientInRoom(client)

		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.encode())
		}
	}
}

func (room *RoomChat) registerClientInRoom(client *Client) {
	if !room.Private {
		room.notifyClientJoined(client)
	}
	room.clients[client] = true
}

func (room *RoomChat) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client]; ok {
		delete(room.clients, client)
	}
}

func (room *RoomChat) broadcastToClientsInRoom(message []byte) {
	for client := range room.clients {
		client.send <- message
	}
}

func (room *RoomChat) notifyClientJoined(client *Client) {
	message := &MessageChat{
		Action:  SendMessageAction,
		Target:  room,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
	}

	room.broadcastToClientsInRoom(message.encode())
}

func (room *RoomChat) GetId() string {
	return room.ID.String()
}

func (room *RoomChat) GetName() string {
	return room.Name
}
