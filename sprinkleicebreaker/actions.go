package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func actionsHandler(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(body))
	payload := c.Request.FormValue("payload")

	// var payloadStruct Payload
	// err := json.Unmarshal([]byte(payload), &payloadStruct)
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"text": "The request failed :(",
	// 	})
	// 	return
	// }
	// if len(payloadStruct.Actions) != 1 {
	// 	c.JSON(200, gin.H{
	// 		"text": "The request failed, expected to receive 1 action :(",
	// 	})
	// 	return
	// }
	// if len(payloadStruct.Actions[0].Value) == 0 {
	// 	fmt.Println(payload)
	// 	c.JSON(200, gin.H{
	// 		"text": "The request failed, expected the action to not be empty :(",
	// 	})
	// 	return
	// }
	actionType := gjson.Get(payload, "actions.0.type")
	if !actionType.Exists() {
		c.JSON(200, gin.H{
			"response_type": "in_channel",
			"text":          "The request failed parsing the action type :(",
		})
		return
	}

	switch actionType.String() {

	case "button":
		buttonValue := gjson.Get(payload, "actions.0.value")
		if !buttonValue.Exists() {
			c.JSON(200, gin.H{
				"response_type": "in_channel",
				"text":          "The request failed parsing the value :(",
			})
			return
		}

		switch buttonValue.String() {

		case "start":
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
			return

		case "cancel":
			c.JSON(200, gin.H{
				"response_type": "in_channel",
				"text":          "Game request cancelled.",
			})
			return
		}

	case "select":
		selectValue := gjson.Get(payload, "actions.0.selected_options.0.value")
		if !selectValue.Exists() {
			c.JSON(200, gin.H{
				"response_type": "in_channel",
				"text":          "The request failed parsing the value :(",
			})
			return
		}

		c.JSON(200, gin.H{
			"response_type": "in_channel",
			"text":          "Your selected option: " + selectValue.String(),
		})
	}

	c.JSON(200, gin.H{
		"response_type": "in_channel",
		"text":          "The request couldn't figure out what to do next :(",
	})
	return
}
