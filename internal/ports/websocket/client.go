package websocket

type IClient interface {
	readMessageFromSocket()
	sendMessageToSocket()
	disconnect()
	handleNewMessage(jsonMessage []byte)
	handleJoinRoomMessage(message MessageChat)
	handleLeaveRoomMessage(message MessageChat)
	handleJoinRoomPrivateMessage(message MessageChat)
	joinRoom(roomName string, sender *Client)
	isInRoom(room *RoomChat) bool
	notifyRoomJoined(room *RoomChat, sender *Client)
	GetName() string
	GetID() string
}
