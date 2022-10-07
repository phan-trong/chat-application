package websocket

import "github.com/google/uuid"

type WebSocketServer interface {
	Run()
	registerClient(client *Client)
	unregisterClient(client *Client)
	broadcastToClients(message []byte)
	notifyClientJoined(client *Client)
	notifyClientLeft(client *Client)
	listOnlineClients(client *Client)
	findRoomByName(roomName string) *RoomChat
	findRoomByID(roomId string) *RoomChat
	createRoom(name string, private bool) *RoomChat
	findClientByID(ID uuid.UUID) *Client
}
