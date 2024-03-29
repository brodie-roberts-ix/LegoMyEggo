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

		case "Let's go!":
			game, err := buildNewGame(gameFilePath)
			if err != nil {
				status200InChannelWithText(c, "Unable to create the escape room adventure :(")
				return
			}
			channelGameState[channelID] = game

			status200WithSelection(c, buttonName)

			desc, actions := game.Display()
			renderLeggoGameReply(channelID, desc, actions)
			return

		case "Not yet.":
			status200InChannelWithText(c, "Come back when you're ready :smile:")
			return

		default:
			status200InChannelWithText(c, "The request failed parsing the button value :(")
			return
		}

	case "select":
		userText := gjson.Get(payload, "actions.0.selected_options.0.value").String()

		game, ok := channelGameState[channelID]
		if !ok {
			status200InChannelWithText(c, "Unable to obtain the escape room state :(")
			return
		}

		status200WithSelection(c, userText)

		msg, actions := game.Act(userText)
		if game.HaveWon() {
			delete(channelGameState, channelID)
			renderLeggoGameMessage(channelID, msg)
			return
		}
		renderLeggoGameReply(channelID, msg, actions)

	default:
		status200InChannelWithText(c, "The request failed parsing the action type :(")
		return
	}
}
