package main

import (
	"github.com/gin-gonic/gin"
)

var integrationURL = "http://fd295619.ngrok.io"

type Action struct {
	Value string `json:"value"`
}
type Payload struct {
	Actions []Action `json:"actions"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/slash-command", slashCommandHandler)
	r.POST("/actions", actionsHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
