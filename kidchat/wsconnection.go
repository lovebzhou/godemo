package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type WSConnection struct {
	c      *Client
	socket *websocket.Conn
}

func (c *WSConnection) read() {
	var c *Connection = wsc
	defer func() {
		manager.unregister <- c.(*Connection)
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		log.Printf("[%s]%s", c.socket.RemoteAddr(), message)
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *WSConnection) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	log.Println("wsPage")
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}

	id, _ := uuid.NewV4()
	wsconn := &WSConnection{id: iid.String(), socket: conn, send: make(chan []byte)}

	manager.register <- wsconn

	go wsconn.read()
	go wsconn.write()
}

func ListenAndServe() {
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(addr, nil)
}
