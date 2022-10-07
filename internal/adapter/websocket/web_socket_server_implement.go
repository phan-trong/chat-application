package websocket

import (
	"chat_application/internal/domain"
	"log"
)

type WsServer struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	rooms      map[*RoomChat]bool
	userRepo   domain.UserRepository
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer(userRepo domain.UserRepository) *WsServer {
	return &WsServer{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		rooms:      make(map[*RoomChat]bool),
		userRepo:   userRepo,
	}
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)

		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *WsServer) registerClient(client *Client) {
	server.notifyClientJoined(client)
	server.listOnlineClients(client)
	server.clients[client.GetID()] = client
}

func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client.GetID()]; ok {
		delete(server.clients, client.GetID())
	}
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		server.clients[client].send <- message
	}
}

func (server *WsServer) notifyClientJoined(client *Client) {
	message := &MessageChat{
		Action: UserJoinedAction,
		Sender: client,
	}

	server.broadcastToClients(message.encode())
}

func (server *WsServer) notifyClientLeft(client *Client) {
	message := &MessageChat{
		Action: UserLeftAction,
		Sender: client,
	}

	server.broadcastToClients(message.encode())
}

func (server *WsServer) listOnlineClients(client *Client) {
	for existingClient := range server.clients {
		message := &MessageChat{
			Action: UserJoinedAction,
			Sender: server.clients[existingClient],
		}
		client.send <- message.encode()
	}
}

func (server *WsServer) findRoomByName(name string) *RoomChat {
	var foundRoom *RoomChat
	for room := range server.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (server *WsServer) findRoomByID(roomId string) *RoomChat {
	var foundRoom *RoomChat
	for room := range server.rooms {
		if room.GetId() == roomId {
			foundRoom = room
			break
		}
	}
	log.Println("foundRoom")
	log.Println(foundRoom)
	return foundRoom
}

func (server *WsServer) createRoom(name string, private bool) *RoomChat {
	room := NewRoom(name, private)
	go room.RunRoom()
	server.rooms[room] = true
	return room
}

func (server *WsServer) findClientByID(uid string) *Client {
	var foundClient *Client
	for client := range server.clients {
		if client == uid {
			foundClient = server.clients[client]
			break
		}
	}
	return foundClient
}
