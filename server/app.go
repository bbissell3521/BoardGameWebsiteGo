package main

/*
This package is designed to serve as a backend framework for
the FreeToGame API. This was designed to be used in conjuction
with a React.JS webpage contained in the same directory.

Author: Blake Bissell
Date: 10/31/2021
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// This is the structure for a single game listing.
type game struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"short_description"`
	GameURL       string `json:"game_url"`
	Genre         string `json:"genre"`
	Platform      string `json:"platform"`
	Publisher     string `json:"publisher"`
	Developer     string `json:"developer"`
	PublishDate   string `json:"release_date"`
	FreeToGameURL string `json:"freetogame_profile_url"`
}

// Test local data for development purposes
var games = []game{
	{
		ID:            "1",
		Title:         "Dauntless",
		Thumbnail:     "https://www.freetogame.com/g/1/thumbnail.jpg",
		Description:   "A free-to-play, co-op action RPG with gameplay similar to Monster Hunter.",
		GameURL:       "https://www.freetogame.com/open/dauntless",
		Genre:         "MMORPG",
		Platform:      "PC (Windows)",
		Publisher:     "Phoenix Labs",
		Developer:     "Phoenix Labs, Iron Galaxy",
		PublishDate:   "2019-05-21",
		FreeToGameURL: "https://www.freetogame.com/dauntless"},
	{
		ID:            "2",
		Title:         "World of Tanks",
		Thumbnail:     "https://www.freetogame.com/g/2/thumbnail.jpg",
		Description:   "If you like blowing up tanks, with a quick and intense game style you will love this game!",
		GameURL:       "https://www.freetogame.com/open/world-of-tanks",
		Genre:         "Shooter",
		Platform:      "PC (Windows)",
		Publisher:     "Wargaming",
		Developer:     "Wargaming",
		PublishDate:   "2011-04-12",
		FreeToGameURL: "https://www.freetogame.com/world-of-tanks"},
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func getGameByID(c *gin.Context) {
	id := c.Param("id")

	for _, game := range games {
		if game.ID == id {
			c.IndentedJSON(http.StatusOK, game)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Game Not Found"})
}

func postGame(c *gin.Context) {
	var newGame game

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGameByID)
	router.POST("/games", postGame)

	router.Run("localhost:8080")
}
