package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"strings"
)

var host = flag.String("h", "localhost", "host")
var port = flag.String("p", "9898", "port")

// Client is a tcp conection and sender buffer
type Client struct {
	socket net.Conn
	data   chan []byte
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			//log.Println("RECEIVED: " + string(message))
			log.Printf("RECEIVED[%s]:%s", client.socket.RemoteAddr(), string(message))
		}
	}
}

// ClientManager is a client manager
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			log.Println("Added new connection!", connection.socket.RemoteAddr())
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				log.Println("A connection has terminated!", connection.socket.RemoteAddr())
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			log.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func startServerMode() {
	addr := *host + ":" + *port
	log.Printf("Starting server... %s", addr)

	listener, error := net.Listen("tcp", addr)
	if error != nil {
		log.Println(error)
	}

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()

	for {
		connection, _ := listener.Accept()
		if error != nil {
			log.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}

func startClientMode() {
	addr := *host + ":" + *port
	log.Println("Starting client...", addr)

	connection, error := net.Dial("tcp", addr)
	if error != nil {
		log.Println(error)
	}

	client := &Client{socket: connection}
	go client.receive()

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServerMode()
	} else {
		startClientMode()
	}
}
