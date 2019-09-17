package main

import (
	"encoding/json"

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
	r.POST("/slash-command", func(c *gin.Context) {
		commandName := c.Request.FormValue("command")
		if commandName != "/sprinkleicebreaker" {
			c.JSON(200, gin.H{
				"text": "Unexpected command: " + commandName,
			})
			return
		}

		commandArgs := c.Request.FormValue("text")
		channelName := c.Request.FormValue("channel_name")
		c.JSON(200, gin.H{
			"response_type": "in_channel",
			"text":          "Welcome! Command name: " + commandName + ", Args: " + commandArgs + ", Channel name: " + channelName,
			"attachments": []gin.H{
				gin.H{
					"fallback":    "This is a fallback for when things didn't work as expected :(",
					"callback_id": "icebreaker",
					"actions":     startIcebreakerButton(),
				},
			},
		})
	})
	r.POST("/actions", func(c *gin.Context) {
		// body, _ := ioutil.ReadAll(c.Request.Body)
		// fmt.Println(string(body))
		payload := c.Request.FormValue("payload")

		var payloadStruct Payload
		err := json.Unmarshal([]byte(payload), &payloadStruct)
		if err != nil {
			c.JSON(404, gin.H{
				"text": "The request failed :(",
			})
		}
		if len(payloadStruct.Actions) != 1 {
			c.JSON(404, gin.H{
				"text": "The request failed, expected to receive 1 action :(",
			})
		}
		if len(payloadStruct.Actions[0].Value) == 0 {
			c.JSON(404, gin.H{
				"text": "The request failed, expected the action to not be empty :(",
			})
		}

		c.JSON(200, gin.H{
			"response_type": "in_channel",
			"text":          "Gotcha! You clicked " + payloadStruct.Actions[0].Value,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
