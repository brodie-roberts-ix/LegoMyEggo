package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func actionsHandler(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(body))
	payload := c.Request.FormValue("payload")

	actionType := gjson.Get(payload, "actions.0.type").String()

	switch actionType {

	case "button":
		buttonValue := gjson.Get(payload, "actions.0.value").String()

		switch buttonValue {

		// case "start_traditional_icebreaker":
		// 	status200InChannelWithText(c, "Question: What is your favorite non-alcoholic beverage?")
		// 	return

		case "start_escape_room":
			status200InChannelWithTextAndMultiSelect(c, "You have arrived in your first room. What do you do?", iceBreakerSelectMenu())
			return

		case "cancel":
			status200InChannelWithText(c, "Game request cancelled.")
			return

		default:
			status200InChannelWithText(c, "The request failed parsing the button value :(")
			return
		}

	case "select":
		selectValue := gjson.Get(payload, "actions.0.selected_options.0.value").String()

		switch selectValue {

		case "option-1":
			fallthrough
		case "option-2":
			message := "Your selected option: " + selectValue + "\nYou are still in the first room. What do you do?"
			status200InChannelWithTextAndMultiSelect(c, message, iceBreakerSelectMenu())
			return

		case "option-3":
			status200InChannelWithText(c, "You found the exit. Congratulations!")
			return

		default:
			status200InChannelWithText(c, "The request failed parsing the select value :(")
			return
		}

	default:
		status200InChannelWithText(c, "The request failed parsing the action type :(")
		return
	}
}
