package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	chatPostMessageURL = "https://slack.com/api/chat.postMessage"
	botOAuthToken      = "xoxb-765348086295-766935288614-x992BfmLUQN50t5pWIJYYhH8"
)

func status200InChannelWithText(c *gin.Context, text string) {
	c.JSON(http.StatusOK, gin.H{
		"response_type": "in_channel",
		"text":          text,
	})
}

func status200InChannelWithTextAndMultiSelect(c *gin.Context, text string, actions []gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"response_type":   "in_channel",
		"text":            text,
		"attachment_type": "default",
		"attachments": []gin.H{
			gin.H{
				"fallback":    "This is a fallback for when things didn't work as expected :(",
				"callback_id": "icebreaker",
				"type":        "static_select",
				"actions":     actions,
			},
		},
	})
}

func postMessageWithText(channelID string, text string) {
	queryParams := url.Values{}
	queryParams.Add("token", botOAuthToken)
	queryParams.Add("channel", channelID)
	queryParams.Add("text", text)

	hc := &http.Client{}

	resp, err := hc.Post(chatPostMessageURL+"?"+queryParams.Encode(), "application/plain-text", strings.NewReader(""))
	if err != nil {
		fmt.Println("There was a problem with sending the additional request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("There was a problem with receiving the additional request")
	}
}

func postMessageMultiSelect(channelID string, text string, actions []gin.H) {
	attachmentStruct := []gin.H{
		gin.H{
			"fallback":    "This is a fallback for when things didn't work as expected :(",
			"callback_id": "icebreaker",
			"type":        "static_select",
			"actions":     actions,
		},
	}

	attachments, _ := json.Marshal(attachmentStruct)
	queryParams := url.Values{}
	queryParams.Add("token", botOAuthToken)
	queryParams.Add("channel", channelID)
	queryParams.Add("text", text)
	queryParams.Add("attachments", string(attachments))

	hc := &http.Client{}

	resp, err := hc.Post(chatPostMessageURL+"?"+queryParams.Encode(), "application/plain-text", strings.NewReader(""))
	if err != nil {
		fmt.Println("There was a problem with sending the additional request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("There was a problem with receiving the additional request")
	}
}

func iceBreakerButtons() []gin.H {
	return []gin.H{
		//button("Start traditional ice-breaker activity", "start_traditional_icebreaker"),
		button("Start escape room activity", "start_escape_room"),
		button("Cancel request", "cancel"),
	}
}
func button(text, value string) gin.H {
	return gin.H{
		"name":  value,
		"text":  text,
		"type":  "button",
		"value": value,
	}
}

func iceBreakerSelectMenu() []gin.H {
	return []gin.H{
		gin.H{
			"name": "icebreaker_options",
			"text": "Here are your options...",
			"type": "select",
			"options": []gin.H{
				option("Option 1", "option-1"),
				option("Option 2", "option-2"),
				option("Option 3", "option-3"),
			},
		},
	}
}
func option(text, value string) gin.H {
	return gin.H{
		"text":  text,
		"value": value,
	}
}
