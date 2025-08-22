package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// // albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func main() {
// 	router := gin.Default()
// 	//getAlbum := 0
// 	router.GET("/albums", getAlbums)

// 	router.Run("localhost:8080")
// }
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)

//}

var quotes = []string{
	"Stay hungry, stay foolish.",
	"Code is like humor. When you have to explain it, it’s bad.",
	"Simplicity is the soul of efficiency.",
	"Before software can be reusable it first has to be usable.",
	"Fix the cause, not the symptom.",
}

func getRandomQuote(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- quotes[rand.Intn(len(quotes))] // Generate a random index and send the quote to the channel

}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	router := gin.Default()

	router.GET("/quote", func(c *gin.Context) {
		ch := make(chan string)
		go getRandomQuote(ch)          // Start a goroutine to get a random quote
		quote := <-ch                  // Wait for the quote to be sent to the channel
		c.String(http.StatusOK, quote) // Send the quote as a response
	})

	router.Run("localhost:8080") // Start the server on port 8080

}
