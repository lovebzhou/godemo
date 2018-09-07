package main

import (
	"log"
)

// Message is
type Message struct {
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Content string `json:"content,omitempty"`
}

// Connection is ...
type Connection interface {
	read()
	write()
}

// Client is
type Client struct {
	id   string
	send chan []byte
	conn Connection
}

// ClientManager is a client manager
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client

	wsPort  string
	tcpPort string
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
	wsPort:     "8080",
	tcpPort:    "5432",
}

func (manager *ClientManager) serve() {
	for {
		select {
		case client := <-manager.register:
			manager.clients[client] = true
		case client := <-manager.unregister:
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
			}
		case message := <-manager.broadcast:
			for client := range manager.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(manager.clients, client)
				}
			}
		}
	}
}

func main() {
	log.Println("starting ...")

	manager.serve()

	log.Println("finish")
}
