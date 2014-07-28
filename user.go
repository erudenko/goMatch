package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type User struct {
	ws     *websocket.Conn
	send   chan []byte
	filter UserFilter
}

// readPump pumps messages from the websocket connection to the hub.
// {"userID":"123"}
func (u *User) readPump() {
	defer func() {
		matcher.unregister <- u
		u.ws.Close()
	}()
	u.ws.SetReadLimit(maxMessageSize)
	u.ws.SetReadDeadline(time.Now().Add(pongWait))
	u.ws.SetPongHandler(func(string) error { u.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := u.ws.ReadMessage()
		if err != nil {
			break
		}
		// log.Println("got message from client:", u, message)
		filter := UserFilter{}
		if err = json.Unmarshal(message, &filter); err != nil || len(filter.UserID) == 0 {
			log.Println("error parsing message:", string(message), err)
			fmt.Printf("wrong formated filter %+v \n", filter)
			break
		}
		u.filter = filter
		log.Println("got client filter:", u, filter)
		matcher.register <- u
		// h.broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (u *User) write(mt int, payload []byte) error {
	u.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return u.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (u *User) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.ws.Close()
	}()
	for {
		select {
		case message, ok := <-u.send:
			if !ok {
				u.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := u.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := u.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveSocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	u := &User{send: make(chan []byte, 256), ws: ws}
	//matcher.register <- u
	go u.writePump()
	u.readPump()
}
