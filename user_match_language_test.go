package main

import (
	"testing"
)

func TestMatchLanguageNotSet(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	mark := user1.matchGender(user2)
	if mark != 2 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageOneSet(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user2 := constructUser2()
	mark := user1.matchLanguage(user2)
	if mark != -1000 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageOneSetReverse(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails1()
	mark := user1.matchLanguage(user2)
	if mark != -1000 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageMatchBoth(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user1.filter.ProvideFilters = *constructFilterDetails2()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails2()
	user2.filter.ProvideFilters = *constructFilterDetails1()
	mark := user1.matchLanguage(user2)
	if mark != 18 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageUnMatchBoth(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters = *constructFilterDetails1()
	user1.filter.ProvideFilters = *constructFilterDetails2()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails1()
	user2.filter.ProvideFilters = *constructFilterDetails2()
	mark := user1.matchLanguage(user2)
	if mark != -1000 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageMatchOneNotSetOther(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.ProvideFilters = *constructFilterDetails1()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails1()
	mark := user1.matchLanguage(user2)
	if mark != 10 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageMatchOneNotSetOtherPartly(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.ProvideFilters = *constructFilterDetails1()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters = *constructFilterDetails3()
	mark := user1.matchLanguage(user2)
	if mark != 6 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchLanguageMatchBothNotSet(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	mark := user1.matchLanguage(user2)
	if mark != 2 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}
