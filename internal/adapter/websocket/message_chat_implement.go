package websocket

import (
	"encoding/json"
	"log"
)

const SendMessageAction = "send-message"
const CreateRoomAction = "create-room"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const JoinRoomPrivateAction = "join-room-private"
const RoomJoinedAction = "room-joined"

type MessageChat struct {
	Action  string    `json:"action"`
	Message string    `json:"message"`
	Target  *RoomChat `json:"target"`
	Sender  *Client   `json:"sender"`
}

func (message *MessageChat) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}

func (message *MessageChat) UnmarshalJSON(data []byte) error {
	type Alias MessageChat
	msg := &struct {
		Sender Client `json:"sender"`
		*Alias
	}{
		Alias: (*Alias)(message),
	}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	message.Sender = &msg.Sender
	return nil
}
