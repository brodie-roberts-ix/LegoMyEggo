package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/brodie-roberts-ix/LegoMyEggo/leggo"
)

func buildNewGame() (*leggo.Game, error) {
	file, err := os.Open("../leggo/stories/derp.json")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	game, err := leggo.NewFromJson(data)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func renderLeggoGameMessage(channelID, message string) {
	go func() {
		time.Sleep(sleepBeforeGameReply)
		postMessageWithText(channelID, message)
	}()
}

func renderLeggoGameReply(channelID, message string, actions []string) {
	go func() {
		time.Sleep(sleepBeforeGameReply)
		postMessageWithText(channelID, message)
		postMessageMultiSelect(channelID, "What do you do?", selectMenu(actions))
	}()
}
