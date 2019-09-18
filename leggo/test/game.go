package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"leggo"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../stories/derp.json")
	tableflip(err)
	data, err := ioutil.ReadAll(file)
	tableflip(err)

	game, err := leggo.NewFromJson(data)
	tableflip(err)

	var msg string
	var actions []string
	msg, actions = game.Display()
	fmt.Println(msg)
	fmt.Println("Actions: {", strings.Join(actions, " | "), "}")

	scanner := bufio.NewScanner(os.Stdin)
	for !game.HaveWon() {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		msg, actions = game.Act(scanner.Text())
		fmt.Println(msg)
		fmt.Println("Actions: {", strings.Join(actions, " | "), "}")
	}
	tableflip(scanner.Err())
	fmt.Println("VICTORY!!!")
}

func tableflip(err error) {
	if err != nil {
		panic(err)
	}
}
