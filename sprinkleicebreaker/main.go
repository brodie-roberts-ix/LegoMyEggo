package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

var integrationURL = "http://fd295619.ngrok.io"

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
					"callback_id": "start_icebreaker",
					"actions": []gin.H{
						gin.H{
							"name":  "start",
							"text":  "Start Icebreaker",
							"type":  "button",
							"value": "start",
						},
					},
				},
			},
		})
	})
	r.POST("/actions", func(c *gin.Context) {
		// body, _ := ioutil.ReadAll(c.Request.Body)
		// fmt.Println(string(body))
		payload := c.Request.FormValue("payload")

		type Action struct {
			value string `json:"value"`
		}
		type Payload struct {
			actions []Action `json:"actions"`
		}
		var payloadStruct Payload
		err := json.Unmarshal([]byte(payload), &payloadStruct)
		if err != nil {
			c.JSON(404, gin.H{
				"text": "The request failed :(",
			})
		}

		c.JSON(200, gin.H{
			"response_type": "in_channel",
			//"text":          "Gotcha! You clicked " + payloadStruct.actions[0].value,
			"text": "Gotcha! You clicked " + fmt.Sprintf("%+v", payloadStruct),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
