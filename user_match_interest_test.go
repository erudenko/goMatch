package main

import (
	"testing"
)

func TestMatchInteresNotSet(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	mark := user1.matchInterests(user2)
	if mark != 0 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchInetrestOneSet(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user1.filter.ProvideFilters = *constructFilterDetails2()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	mark := user1.matchInterests(user2)
	if mark != 0 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchInetrestFullMatch(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user1.filter.ProvideFilters = *constructFilterDetails2()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails1()
	user2.filter.ProvideFilters = *constructFilterDetails2()
	mark := user1.matchInterests(user2)
	if mark != 6 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchInetrestFullMismatch(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user1.filter.ProvideFilters = *constructFilterDetails2()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails1()
	user2.filter.ProvideFilters = *constructFilterDetails3()
	mark := user1.matchInterests(user2)
	if mark != 0 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}
