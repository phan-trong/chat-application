package websocket

import (
	"chat_application/internal/domain"
	"chat_application/internal/ports/middlewares"
	"chat_application/internal/usecase/message"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	conn        *websocket.Conn
	wsServer    *WsServer
	send        chan []byte
	rooms       map[*RoomChat]bool
	messageRepo message.Repository
}

func newClient(conn *websocket.Conn, wsServer *WsServer, messageRepo message.Repository, userId uuid.UUID, fullName string) *Client {
	return &Client{
		ID:          userId,
		Name:        fullName,
		conn:        conn,
		wsServer:    wsServer,
		send:        make(chan []byte, 256),
		rooms:       make(map[*RoomChat]bool),
		messageRepo: messageRepo,
	}
}

func (client *Client) readMessageFromSocket() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		fmt.Println("Message Receiving: ", string(jsonMessage[:]))

		client.handleNewMessage(jsonMessage)
	}
}

func (client *Client) sendMessageToSocket() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			fmt.Println("Message Sending: ", string(message[:]))

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	for room := range client.rooms {
		room.unregister <- client
	}
	close(client.send)
	client.conn.Close()
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(wsServer *WsServer, c *gin.Context, mRepo message.Repository) {
	userData := middlewares.GetCurrentUser(c)
	userId, ok := c.Request.URL.Query()["id"]

	if !ok || len(userId[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer, mRepo, userData.ID, userData.FullName)
	fmt.Println("New Client joined the hub!")
	fmt.Println(client)
	// Receive Message from socket
	go client.readMessageFromSocket()
	// Send Message to socket
	go client.sendMessageToSocket()

	wsServer.register <- client
}

func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message MessageChat
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}

	message.Sender = client

	switch message.Action {
	case SendMessageAction:

		roomID := message.Target.GetId()
		newMsg, err := domain.NewMessage(message.Sender.ID, message.Target.ID, message.Message)
		if err != nil {
			log.Printf("Can not Create New Message %s", err)
			return
		}
		client.messageRepo.Create(newMsg)
		if err != nil {
			log.Printf("Can not Save New Message %s", err)
			return
		}
		if room := client.wsServer.findRoomByID(roomID); room != nil {
			room.broadcast <- &message
		}

	case JoinRoomAction:
		client.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)

	case JoinRoomPrivateAction:
		client.handleJoinRoomPrivateMessage(message)
	}
}

func (client *Client) handleJoinRoomMessage(message MessageChat) {
	roomName := message.Message

	client.joinRoom(roomName, nil)
}

func (client *Client) handleLeaveRoomMessage(message MessageChat) {
	room := client.wsServer.findRoomByID(message.Message)
	if room == nil {
		return
	}

	if _, ok := client.rooms[room]; ok {
		delete(client.rooms, room)
	}

	room.unregister <- client
}

func (client *Client) handleJoinRoomPrivateMessage(message MessageChat) {
	target := client.wsServer.findClientByID(message.Message)

	if target == nil {
		return
	}

	roomNameArr := []string{target.GetName(), client.GetName()}

	sort.Strings(roomNameArr)
	roomName := strings.Join(roomNameArr, "_")

	client.joinRoom(roomName, target)
	target.joinRoom(roomName, client)
}

func (client *Client) joinRoom(roomName string, sender *Client) {
	room := client.wsServer.findRoomByName(roomName)
	if room == nil {
		room = client.wsServer.createRoom(roomName, sender != nil)
		messageCreateRoom := &MessageChat{
			Action: CreateRoomAction,
			Target: room,
		}
		for clientOnlineId := range client.wsServer.clients {
			if client.ID.String() != clientOnlineId {
				client.wsServer.clients[clientOnlineId].send <- messageCreateRoom.encode()
			}
		}
	}

	// Don't allow to join private rooms through public room message
	if sender == nil && room.Private {
		return
	}

	if !client.isInRoom(room) {

		client.rooms[room] = true
		room.register <- client

		client.notifyRoomJoined(room, sender)
	}
}

func (client *Client) isInRoom(room *RoomChat) bool {
	if _, ok := client.rooms[room]; ok {
		return true
	}

	return false
}

func (client *Client) notifyRoomJoined(room *RoomChat, sender *Client) {
	message := MessageChat{
		Action: RoomJoinedAction,
		Target: room,
		Sender: sender,
	}

	client.send <- message.encode()
}

func (client *Client) GetName() string {
	return client.Name
}

func (client *Client) GetID() string {
	return client.ID.String()
}
