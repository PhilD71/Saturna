// Simple little REST API to test deployment in Kubernetes
// Phillip Dumitru Jul 18th 2023
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Message string `json:msg`
}

func main() {
	router := gin.Default()
	router.GET("/hello", getHello)
	router.POST("/echo", postEcho)
	router.Run("0.0.0.0:8080")
}

// Simple get hello function
// Returns a greeting message
func getHello(c *gin.Context) {
	var helloMsg message = message{Message: "hello from port 8080!"}

	c.IndentedJSON(http.StatusOK, helloMsg)
}

// A "fake" post method, "creates" an entry for the message you sent
// Really, it just replays the message you sent it back
func postEcho(c *gin.Context) {
	var newMsg message

	if err := c.BindJSON(&newMsg); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newMsg)
}
