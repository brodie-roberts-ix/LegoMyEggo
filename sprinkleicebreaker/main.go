package main

import (
	"time"

	"github.com/brodie-roberts-ix/LegoMyEggo/leggo"
	"github.com/gin-gonic/gin"
)

var (
	chatPostMessageURL   = "https://slack.com/api/chat.postMessage"
	botOAuthToken        = "xoxb-765348086295-766935288614-YM0uoAzUma1uGkJvzIoegBL0"
	channelGameState     = make(map[string]*leggo.Game)
	sleepBeforeGameReply = 600 * time.Millisecond
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/slash-command", slashCommandHandler)
	r.POST("/actions", actionsHandler)
	r.POST("/events", eventsHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
