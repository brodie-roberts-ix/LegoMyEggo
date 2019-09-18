package main

import (
	"fmt"

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

	fmt.Println(gjson.Get(payload, "actions.0").String())

	actionType := gjson.Get(payload, "actions.0.type").String()

	switch actionType {

	case "button":
		buttonName := gjson.Get(payload, "actions.0.name").String()
		buttonValue := gjson.Get(payload, "actions.0.value").String()

		switch buttonValue {

		case "Start escape room activity":
			game, err := buildNewGame()
			if err != nil {
				status200InChannelWithText(c, "Unable to create the escape room adventure :(")
				return
			}
			channelGameState[channelID] = game

			status200WithSelection(c, buttonName)

			desc, actions := game.Display()
			renderLeggoGameReply(c, channelID, desc, actions)
			return

		case "Start debug flow":
			status200WithSelection(c, buttonName)
			go func() {
				postMessageWithText(channelID, "The (debug) adventure begins! You have arrived in your first room.")
				postMessageMultiSelect(channelID, "What do you do?", iceBreakerSelectMenu())
			}()
			return

		case "Cancel request":
			status200InChannelWithText(c, "Game request cancelled.")
			return

		default:
			status200InChannelWithText(c, "The request failed parsing the button value :(")
			return
		}

	case "select":
		selectValue := gjson.Get(payload, "actions.0.selected_options.0.value").String()

		switch selectValue {

		case "Option 1":
			fallthrough
		case "Option 2":
			status200WithSelection(c, selectValue)
			go func() {
				postMessageWithText(channelID, "You are still in the first room.")
				postMessageMultiSelect(channelID, "What do you do?", iceBreakerSelectMenu())
			}()
			return

		case "Option 3":
			status200WithSelection(c, selectValue)
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
