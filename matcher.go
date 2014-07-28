package main

import (
	"fmt"
	"log"
	"time"
)

type Matcher struct {
	connections map[*User]bool
	register    chan *User
	unregister  chan *User
}

var matcher = Matcher{
	register:    make(chan *User),
	unregister:  make(chan *User),
	connections: make(map[*User]bool),
}

var connectMessage = "{\"userID\":\"%s\",\"message\":\"connected\"}"

func (m *Matcher) run() {
	ticker := time.NewTicker(time.Millisecond * 500)
	connectionsChanged := false
	for {
		select {
		case u := <-m.register:
			fmt.Printf("+++registered user: %+v\n", u)
			m.connections[u] = true
			connectionsChanged = true
			u.send <- []byte(fmt.Sprintf(connectMessage, u.filter.UserID))
		case u := <-m.unregister:
			log.Println("---unregistered  user: ", u)
			delete(m.connections, u)
			close(u.send)
		case <-ticker.C:
			if connectionsChanged {
				m.match()
				connectionsChanged = false
			}
		}
	}

}

func (m *Matcher) match() {
	//Do nothing yer

}
