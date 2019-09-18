package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func actionsHandler(c *gin.Context) {
	payload := c.Request.FormValue("payload")

	channelID := gjson.Get(payload, "channel.id").String()
	if len(channelID) == 0 {
		status200InChannelWithText(c, "Error: unable to get the channel ID from the request")
		return
	}

	actionType := gjson.Get(payload, "actions.0.type").String()

	switch actionType {

	case "button":
		buttonValue := gjson.Get(payload, "actions.0.value").String()

		switch buttonValue {

		// NOTE: Abandoning idea of traditional ice breaker
		// case "start_traditional_icebreaker":
		// 	status200InChannelWithText(c, "Question: What is your favorite non-alcoholic beverage?")
		// 	return

		case "start_escape_room":
			status200InChannelWithText(c, "Not implemented yet")
			return

		case "start_debug_flow":
			status200InChannelWithText(c, "You have selected \"Start Debug Flow\"")
			go func() {
				postMessageWithText(channelID, "The (debug) adventure begins!")
				postMessageMultiSelect(channelID, "You have arrived in your first room. What do you do?", iceBreakerSelectMenu())
			}()
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
			status200InChannelWithText(c, "You have selected \""+selectValue+"\"")
			go func() {
				postMessageWithText(channelID, "You are still in the first room.")
				postMessageMultiSelect(channelID, "What do you do?", iceBreakerSelectMenu())
			}()
			return

		case "option-3":
			status200InChannelWithText(c, "You have selected \""+selectValue+"\"")
			go func() {
				postMessageWithText(channelID, "You found the exit. Congratulations!")
			}()
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
