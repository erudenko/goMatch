package main

type UserFilter struct {
	provideFilters []string `json:"provideFilters"`
	requestFilters []string `json:"requestFilters"`
	userID         string   `json:"userID"`
}
