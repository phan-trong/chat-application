package domain

import "time"

/*
 *	If message Peer to Peer set SenderId and ReceiverId
 *	If message send to channel set SenderId and RoomId
 */

type Message struct {
	ID         int
	SenderId   int
	ReceiverId int
	RoomId     int
	Message    string
	DeletedAt  time.Time
	CreateAt   time.Time
	UpdatedAt  time.Time
}
