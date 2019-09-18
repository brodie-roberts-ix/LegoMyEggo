package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/brodie-roberts-ix/LegoMyEggo/leggo"
	"github.com/gin-gonic/gin"
)

var (
	chatPostMessageURL   = "https://slack.com/api/chat.postMessage"
	botOAuthToken        = "" // Taken in from the first argument after the program name
	channelGameState     = make(map[string]*leggo.Game)
	sleepBeforeGameReply = 600 * time.Millisecond
	gameFilePath         = "../leggo/stories/derp.json"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Please pass in the app bot's OAuth token.")
		return
	}

	botOAuthToken = flag.Arg(0)
	if len(botOAuthToken) == 0 {
		fmt.Println("There was a problem handling the bot's OAuth token.")
		return
	}

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
