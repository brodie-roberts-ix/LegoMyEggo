package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func slashCommandHandler(c *gin.Context) {
	commandName := c.Request.FormValue("command")
	if commandName != "/sprinkle" {
		c.JSON(http.StatusOK, gin.H{
			"text": "Unexpected command: " + commandName,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response_type": "in_channel",
		"text":          "Are you ready for an adventure?",
		"attachments": []gin.H{
			gin.H{
				"fallback":    "This is a fallback for when things didn't work as expected :(",
				"callback_id": "icebreaker",
				"actions":     iceBreakerButtons(),
			},
		},
	})
}
