package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "reflect"
	"io/ioutil"
	"encoding/json"
)

type Result struct {
	Data []Player `json:"data"`
}

type Player struct {
	PlayerID string `json:"playerId"`
	Name string `json:"name"`
	JerseyNumber int `json:"jerseyNumber"`
}

func main() {
	router := gin.Default()
	router.GET("/:id", getPlayer) // Handle GET requests to "/"
	router.Run(":5000")
}

func getPlayer(c *gin.Context) {
	// Grab the result of our GET request
	result := apiReq()

	players := result.Data
	// Iterate through players and find the player with the correct ID
	for i := 0; i < len(players); i++ {
		if players[i].PlayerID == c.Param("id") {
			// Respond with JSON containing player name and jersey number
			c.JSON(http.StatusOK, gin.H{
				"name": players[i].Name,
				"jerseyNumber": players[i].JerseyNumber,
			})
		}
	}
}

func apiReq() Result {
	// Make GET request to iSports API
	url := "http://api.isportsapi.com/sport/basketball/player?api_key=5Wg8xgsxftGDBBZ6"
	response, err := http.Get(url)

	// If we received an error, print it and respond with an error
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close() // Close the response body after this function's execution

	// Get data and string from the response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	responseString := string(responseData)

	// Convert to player
	var result Result
	json.Unmarshal([]byte(responseString), &result)

	return result;
}

