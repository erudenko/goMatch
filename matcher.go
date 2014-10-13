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
	ticker := time.NewTicker(time.Millisecond * 1000) //match every second
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
	//Do nothing yet
	matchResults := make(map[*User]*User)
	for user, registered := range m.connections {
		//if we already match users with the criterias, skip it
		_, exists := matchResults[user]
		submatchResults := make(map[*User]int)
		if !exists && registered {
			for user2, registered := range m.connections {
				if user2 != user && registered {
					submatchResults[user2] = user.match(user2)
				}
			}
		}
	}

}
