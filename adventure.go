package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/brodie-roberts-ix/LegoMyEggo/leggo"
)

var storyFile = "stories/demo.json"

func main() {
	if len(os.Args) > 1 {
		storyFile = os.Args[1]
	}

	file, err := os.Open(storyFile)
	tableflip(err)
	data, err := ioutil.ReadAll(file)
	tableflip(err)

	game, err := leggo.NewFromJson(data)
	tableflip(err)
	fmt.Printf("Loading game file: %s\n", storyFile)

	var msg string
	var actions []string
	msg, actions = game.Display()
	fmt.Println()
	fmt.Println(msg)
	fmt.Println("  {", strings.Join(actions, " | "), "}\n")

	scanner := bufio.NewScanner(os.Stdin)
	for !game.HaveWon() {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		msg, actions = game.Act(line)
		fmt.Println()
		fmt.Println(msg)
		fmt.Println("  {", strings.Join(actions, " | "), "}\n")
	}
	tableflip(scanner.Err())

	if game.HaveWon() {
		fmt.Println("VICTORY!!!")
	} else {
		fmt.Println("Goodbye.")
	}
}

func tableflip(err error) {
	if err != nil {
		panic(err)
	}
}
