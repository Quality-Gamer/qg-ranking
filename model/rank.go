package model

import (
	"github.com/dayvson/go-leaderboard"
	"os"
)

func GetRank() leaderboard.Leaderboard {
	rank := leaderboard.NewLeaderboard(leaderboard.RedisSettings{
		Host: os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
	}, "gm:rr", 5)

	return rank
}
