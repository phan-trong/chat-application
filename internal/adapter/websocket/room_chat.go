package websocket

type IRoomChat interface {
	RunRoom()
	registerClientInRoom(client *Client)
	unregisterClientInRoom(client *Client)
	broadcastToClientsInRoom(message []byte)
	notifyClientJoined(client *Client)
	GetId() string
	GetName() string
}
