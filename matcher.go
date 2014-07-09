package main

import (
	"time"
)

type Matcher struct {
	users map[*User]bool

	register   chan *User
	unregister chan *User
}

type matchPair struct {
	user1, user2 *User
}

var sharedMatcher = Matcher{
	users:      make(map[*User]bool),
	register:   make(chan *User),
	unregister: make(chan *User),
}

func (matcher *Matcher) runLoop() {
	go matcher.match()
	for {
		select {
		case user := <-matcher.register:
			matcher[user] = true
		case user := <-matcher.unregister:
			user.dropConnection()
			delete(matcher.users, user)
		}
	}
}

func newMatchPair(user1, user2 *User) *matchPair {
	var first *User
	var second *User
	if (user1.userID>user2.iserID) {
		first = user1
		second = user2
	} else {
		first = user2
		second = user1
	}
	return &User(user1:first, user2:second)
}

func (matcher *Matcher) match() {
	//matching map
	matchingHistory := make(map[matchPair]int)
	for {
		//match result is match one user to another with match points
		matchResult := make(map[matchPair]int)
		//make copy of current map
	}
}
