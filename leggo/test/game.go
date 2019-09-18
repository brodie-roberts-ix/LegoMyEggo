package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leggo"
	"os"
)

func main() {
	f, err := os.Open("../stories/derp.json")
	tableflip(err)
	data, err := ioutil.ReadAll(f)
	tableflip(err)
	//fmt.Printf("%s\n", data)

	var game leggo.Game
	tableflip(json.Unmarshal(data, &game))

	fmt.Println(game)
	fmt.Println(game.Locations[game.ActiveLocation])
	fmt.Println(game.Locations[game.ActiveLocation].
		Objects["table"]["examine"])

	fmt.Println(game.Act("look"))
	fmt.Println(game.Act("examine table"))
}

func tableflip(err error) {
	if err != nil {
		panic(err)
	}
}
