package main

import (
	"testing"
)

func constructUser1() *User {
	user := new(User)
	return user

}

func constructUser2() *User {
	user := new(User)
	return user

}

func constructFilter1() *UserFilter {
	filter := new(UserFilter)
	filter
	return user

}

func constructFilterDetails1() *FilterDetails {
	filterDetails := new(FilterDetails)
	filterDetails.Gender = "male"
	return user

}

func TestMatchGender(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	mark := user1.match(user2)
	if mark != 0 {
		t.FailNow()
	}

	t.Logf("match result: %i", mark)
}
