package game

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRoomID() string {
	b := make([]byte, 6)
	for i := range b {
		randomByte, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[randomByte.Int64()]
	}
	return string(b)
}

type Room struct {
	id      string
	players []*Client
	game    *Game
	hub     *Hub
}

func NewRoom(hub *Hub) *Room {
	roomID := generateRoomID()
	// for room in hub.rooms
	if hub.rooms[roomID] != nil {
		roomID = generateRoomID()
	}

	return &Room{
		id:      roomID,
		players: make([]*Client, 0),
		hub:     hub,
	}
}
