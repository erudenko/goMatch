package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

type Matcher struct {
	connections map[*User]bool
	register    chan *User
	unregister  chan *User
	log         map[*MatchLogEntity]int
}

type MatchLogEntity struct {
	UserID1 string
	UserID2 string
}

type MatchEntity struct {
	owner *User
	user  *User
	score int
}

type ByScore []*MatchEntity

var matcher = Matcher{
	register:    make(chan *User),
	unregister:  make(chan *User),
	connections: make(map[*User]bool),
	log:         make(map[*MatchLogEntity]int),
}

var matchMessage = "{\"type\":\"system\",\"id\":\"%s\",\"message\":\"connect_to\",\"connect_id\":\"%s\"}"
var connectMessage = "{\"userID\":\"%s\",\"message\":\"connected\"}"

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score < a[j].score }

func (l *MatchLogEntity) fillUsers(user1 *User, user2 *User) {
	if user1.filter.UserID > user2.filter.UserID {
		l.UserID1 = user1.filter.UserID
		l.UserID2 = user2.filter.UserID
	} else {
		l.UserID1 = user2.filter.UserID
		l.UserID2 = user1.filter.UserID
	}
}

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
	maxMatchScores := make(map[*User]*MatchEntity)
	for user, registered := range m.connections {
		//if we already match users with the criterias, skip it
		_, exists := matchResults[user]
		submatchResults := make(map[*User]int)
		if !exists && registered {
			for user2, registered := range m.connections {
				if (user2 != user) && registered {
					//-10 to match if users were matched before
					//multiply on time of matching
					log := MatchLogEntity{}
					log.fillUsers(user, user2)
					matchBeforePenalty := m.log[&log]
					score := user.match(user2) - (10 * matchBeforePenalty) //we have all matched results for thsi user
					if score >= 0 {
						submatchResults[user2] = score
					}
				}
			}
		}
		//we have matchings
		if len(submatchResults) > 0 {
			matchArray := make([]*MatchEntity, len(submatchResults))
			var counter = 0
			for k, v := range submatchResults {
				matchEntity := MatchEntity{user: k, score: v}
				matchArray[counter] = &matchEntity
				counter++
			}
			sort.Sort(sort.Reverse(ByScore(matchArray)))
			resultBestMatch := matchArray[0]
			resultBestMatch.owner = user
			maxMatchScores[user] = resultBestMatch
			matchResults[user] = resultBestMatch.user
		}

		//sort matches
		//1. build field already matched users
		//2. sort,  decreasing priority of users, that was matched before
		//3. build matching pairs list using priorities
		//4. return result for matching users
		//5. disconnect that users
	}
	//now sort matches accortding to max match score
	if len(maxMatchScores) > 0 {
		maxMatchArray := make([]*MatchEntity, len(maxMatchScores))
		var counter = 0
		for _, v := range maxMatchScores {
			maxMatchArray[counter] = v
			counter++
		}
		//go throught array and send match results
		matchResultsSend := make(map[*User]bool)
		sort.Sort(ByScore(maxMatchArray))
		for _, match := range maxMatchArray {
			if !matchResultsSend[match.owner] {
				m.sendMatchResult(match.owner, match.user)
				matchResultsSend[match.owner] = true
			}
			if !matchResultsSend[match.user] {
				m.sendMatchResult(match.user, match.owner)
				matchResultsSend[match.user] = true
			}
		}
	}
}

func (m *Matcher) sendMatchResult(userToSend *User, userMatched *User) {
	// {"type":"system","id":"USER_ID","message":"connect_to","connect_id":"TO_USER_ID"}
	userToSend.send <- []byte(fmt.Sprintf(matchMessage, userToSend.filter.UserID, userMatched.filter.UserID))
	delete(m.connections, userToSend)
	close(userToSend.send)
}
