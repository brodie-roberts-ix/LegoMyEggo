package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func status200InChannelWithText(c *gin.Context, text string) {
	c.JSON(http.StatusOK, gin.H{
		"response_type": "in_channel",
		"text":          text,
	})
}

func status200WithSelection(c *gin.Context, selection string) {
	c.JSON(http.StatusOK, gin.H{
		"response_type": "in_channel",
		"text":          "> You have selected: *" + selection + "*",
	})
}

// NOTE: Unused for now
/*
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
*/

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
		button("Let's go!"),
		button("Not yet."),
	}
}
func button(text string) gin.H {
	return gin.H{
		"name":  text,
		"text":  text,
		"type":  "button",
		"value": text,
	}
}

func selectMenu(options []string) []gin.H {
	optionsStruct := make([]gin.H, 0)
	for _, opt := range options {
		optionsStruct = append(optionsStruct, createOption(opt))
	}

	return []gin.H{
		gin.H{
			"name":    "icebreaker_options",
			"text":    "Here are your options...",
			"type":    "select",
			"options": optionsStruct,
		},
	}
}
func createOption(text string) gin.H {
	return gin.H{
		"text":  text,
		"value": text,
	}
}
