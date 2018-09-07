package main

import (
	"log"
	"net"
)

// TCPConnection is ...
type TCPConnection struct {
	socket net.Conn
}

func (c *TCPConnection) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		message := make([]byte, 4096)
		length, err := c.socket.Read(message)
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		if length > 0 {
			log.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (c *TCPConnection) write() {
	defer c.socket.Close()
	for {
		select {
		case message, ok := <-c.data:
			if !ok {
				return
			}
			c.socket.Write(message)
		}
	}
}

func startServerMode() {
	addr := manager.tcpPort
	log.Printf("Starting server... %s", addr)

	listener, error := net.Listen("tcp", addr)
	if error != nil {
		log.Println(error)
	}

	for {
		socket, _ := listener.Accept()
		if error != nil {
			log.Println(error)
		}
		tcpConn := &TCPConnection{socket: socket, data: make(chan []byte)}
		manager.register <- tcpConn
		go manager.receive(tcpConn)
		go manager.send(tcpConn)
	}
}
