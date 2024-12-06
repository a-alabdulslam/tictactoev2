package game

import (
	"fmt"
)

// Hub is a struct that holds all the clients and the messages that are sent to them
type Hub struct {
	// Registered rooms.
	rooms map[string]map[*Client]bool
	// clients
	clients map[string]*Client
	//unregister client.
	leave chan *Client
	//register client.
	login chan *Client
	// Inbound messages from the clients.
	broadcast chan Message
}

// Message struct to hold message data
type Message struct {
	Type      string `json:"type"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
	ID        string `json:"id"`
}

func NewHub() *Hub {
	return &Hub{
		rooms:     make(map[string]map[*Client]bool),
		login:     make(chan *Client),
		leave:     make(chan *Client),
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
	}
}

// Core function to run the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.login:
			h.Login(client)
		case client := <-h.leave:
			h.RemoveClient(client)
			// Broadcast a message to all clients.
		case message := <-h.broadcast:

			//Check if the message is a type of "message"
			h.HandleMessage(message)

		}
	}
}

// function to add client to hub
func (h *Hub) Login(client *Client) {
	h.clients[client.ID] = client
	client.send <- Message{Type: "systemMessage", Sender: "system", Recipient: client.ID, Content: "You are now logged in", ID: "system"}

}
func (h *Hub) CreateRoom(client *Client) {
	room := NewRoom(h)
	h.rooms[room.id] = make(map[*Client]bool)
	h.JoinRoom(client, room.id)
	fmt.Println("Room created: ", room.id)
}

// function check if room exists and if not create it and add client to it
func (h *Hub) JoinRoom(client *Client, roomID string) {
	connections := h.rooms[roomID]

	if connections == nil {
		//close(client.send)
		fmt.Println("Room does not exist")
		//client.send <- Message{Type: "message", Content: "Room does not exist", ID: "system"}
		return
	}

	connections[client] = true
	client.RoomID = roomID
	client.send <- Message{Type: "message", Content: "You are now connected to room " + client.RoomID, ID: "system"}
	fmt.Println("Size of room: ", len(connections))
}

// function to remvoe client from room
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.rooms[client.RoomID]; ok {
		delete(h.rooms[client.RoomID], client)
		close(client.send)
		fmt.Println("Player removed from room: ", client.RoomID)
	}
}

// function to handle message based on type of message
func (h *Hub) HandleMessage(message Message) {
	if message.Type == "systemMessage" {

		client := h.clients[message.Sender]
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.rooms[message.ID], client)
		}
	}
	//Check if the message is a type of "message"
	if message.Type == "message" {
		clients := h.rooms[message.ID]
		fmt.Println(clients)
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.rooms[message.ID], client)
			}
		}
	}

	//Check if the message is a type of "message"
	if message.Type == "newRoom" {
		fmt.Println("Creating new room")

		client := h.clients[message.Sender]

		if client == nil {
			return
		}

		h.CreateRoom(client)

	}

	if message.Type == "joinRoom" {
		client := h.clients[message.Sender]

		if client == nil {
			return
		}
		h.JoinRoom(client, message.ID)

	}

}
