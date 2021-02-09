package main

import (
	"github.com/labstack/echo"
	"os"
	"github.com/Quality-Gamer/qg-ranking/endpoint"
)

//main contains all API endpoints
func main() {
	e := echo.New()

	//Post/Rank
	e.POST("/api/post/rank", endpoint.PostRank)

	//Get/Rank
	e.GET("/api/get/rank", endpoint.GetRank)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}