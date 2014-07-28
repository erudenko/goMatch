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
	return filter

}

func constructFilterDetails1() *FilterDetails {
	filterDetails := new(FilterDetails)
	filterDetails.Gender = "male"
	filterDetails.CountryCodes = []string{"ru", "ua", "en"}
	filterDetails.LanguageCodes = []string{"au", "nz", "ua"}
	filterDetails.Interests = []string{"business", "fashion"}
	return filterDetails
}

func constructFilterDetails2() *FilterDetails {
	filterDetails := new(FilterDetails)
	filterDetails.Gender = "male"
	filterDetails.CountryCodes = []string{"ru", "ua", "en"}
	filterDetails.LanguageCodes = []string{"au", "nz", "ua"}
	filterDetails.Interests = []string{"business", "fashion"}
	return filterDetails
}

func constructFilterDetails3() *FilterDetails {
	filterDetails := new(FilterDetails)
	filterDetails.Gender = "male"
	filterDetails.CountryCodes = []string{"ru", "ua", "en"}
	filterDetails.LanguageCodes = []string{"au", "nz", "ua"}
	filterDetails.Interests = []string{"business", "fashion"}
	return filterDetails
}

func constructFilterDetails4() *FilterDetails {
	filterDetails := new(FilterDetails)
	filterDetails.Gender = "male"
	filterDetails.CountryCodes = []string{"ru", "ua", "en"}
	filterDetails.LanguageCodes = []string{"au", "nz", "ua"}
	filterDetails.Interests = []string{"business", "fashion"}
	return filterDetails
}

func TestMatchGenderNotSet(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	mark := user1.matchGender(user2)
	if mark != 2 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderOneSet(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters.Gender = "male"
	user2 := constructUser2()
	mark := user1.matchGender(user2)
	if mark != -1000 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderOneSetReverse(t *testing.T) {
	user1 := constructUser1()
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters.Gender = "male"

	mark := user1.matchGender(user2)
	if mark != -1000 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderOneSetAndMatch(t *testing.T) {
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.RequestFilters.Gender = "male"
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.ProvideFilters.Gender = "male"
	mark := user1.matchGender(user2)
	if mark != 6 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderOneSetAndMatchReverse(t *testing.T) {
	user2 := constructUser1()
	user2.filter = *new(UserFilter)
	user2.filter.RequestFilters.Gender = "male"
	user1 := constructUser2()
	user1.filter = *new(UserFilter)
	user1.filter.ProvideFilters.Gender = "male"
	mark := user1.matchGender(user2)
	if mark != 6 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderBothSet(t *testing.T) {
	user1 := constructUser2()
	user1.filter = *new(UserFilter)
	user1.filter.ProvideFilters.Gender = "male"
	user1.filter.RequestFilters.Gender = "female"
	user2 := constructUser1()
	user2.filter = *new(UserFilter)
	user2.filter.ProvideFilters.Gender = "female"
	user2.filter.RequestFilters.Gender = "male"
	mark := user1.matchGender(user2)
	if mark != 10 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}

func TestMatchGenderBothSetReverse(t *testing.T) {
	user2 := constructUser2()
	user2.filter = *new(UserFilter)
	user2.filter.ProvideFilters.Gender = "male"
	user2.filter.RequestFilters.Gender = "female"
	user1 := constructUser1()
	user1.filter = *new(UserFilter)
	user1.filter.ProvideFilters.Gender = "female"
	user1.filter.RequestFilters.Gender = "male"
	mark := user1.matchGender(user2)
	if mark != 10 {
		t.Logf("match result: %i", mark)
		t.FailNow()
	}
}
