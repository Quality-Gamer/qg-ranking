package model

import "github.com/dayvson/go-leaderboard"

type Response struct {
	Status   string `json:"status"`
	Response Res    `json:"response"`
	Message  string `json:"message"`
}

type Res struct {
	Rank  []leaderboard.User `json:"rank"`
	Total int            `json:"total"`
}
