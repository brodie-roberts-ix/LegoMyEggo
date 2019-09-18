package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func actionsHandler(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(body))
	payload := c.Request.FormValue("payload")

	var payloadStruct Payload
	err := json.Unmarshal([]byte(payload), &payloadStruct)
	if err != nil {
		c.JSON(200, gin.H{
			"text": "The request failed :(",
		})
		return
	}
	if len(payloadStruct.Actions) != 1 {
		c.JSON(200, gin.H{
			"text": "The request failed, expected to receive 1 action :(",
		})
		return
	}
	if len(payloadStruct.Actions[0].Value) == 0 {
		fmt.Println(payload)
		c.JSON(200, gin.H{
			"text": "The request failed, expected the action to not be empty :(",
		})
		return
	}

	// c.JSON(200, gin.H{
	// 	"response_type": "in_channel",
	// 	"text":          "Gotcha! You clicked " + payloadStruct.Actions[0].Value,
	// })
	c.JSON(200, gin.H{
		"response_type":   "in_channel",
		"text":            "Welcome!",
		"attachment_type": "default",
		"attachments": []gin.H{
			gin.H{
				"fallback":    "This is a fallback for when things didn't work as expected :(",
				"callback_id": "icebreaker",
				"type":        "static_select",
				"actions":     iceBreakerSelectMenu(),
			},
		},
	})
}
