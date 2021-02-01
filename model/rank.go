package model

import "github.com/dayvson/go-leaderboard"

func GetRank() leaderboard.Leaderboard {
	rank := leaderboard.NewLeaderboard(leaderboard.RedisSettings{
		Host: "localhost:6379",
		Password: "",
	}, "gm:rr", 5)

	return rank
}
