package main

import (
	"github.com/gin-gonic/gin"
)

func status200InChannelWithText(c *gin.Context, text string) {
	c.JSON(200, gin.H{
		"response_type": "in_channel",
		"text":          text,
	})
}

func status200InChannelWithTextAndMultiSelect(c *gin.Context, text string, actions []gin.H) {
	c.JSON(200, gin.H{
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

// NOTE: Unused for the time being
/*
func sendAdditionalMessageMultiSelect(responseURL string, text string, actions []gin.H) {
	newMessage := gin.H{
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
	}

	bytes, _ := json.Marshal(newMessage)
	payload := strings.NewReader(string(bytes))

	hc := &http.Client{}

	resp, err := hc.Post(responseURL, "application/json", payload)
	if err != nil {
		fmt.Println("There was a problem with sending the additional request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("There was a problem with receiving the additional request")
	}
}
*/

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
