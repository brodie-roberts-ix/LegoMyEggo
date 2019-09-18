package leggo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Game struct {
	ActiveActions  []string             `json:"start_actions"`
	Actions        map[string]*Action   `json:"actions"`
	ActiveLocation string               `json:"start_location"`
	Locations      map[string]*Location `json:"locations"`
	victory        bool
}

func NewFromJson(data []byte) (*Game, error) {
	var game Game
	err := json.Unmarshal(data, &game)
	return &game, err
}

func (g *Game) HaveWon() bool {
	return g.victory
}

// Returns the current state of the game: room description and actions
func (g *Game) Display() (desc string, actions []string) {
	return g.Locations[g.ActiveLocation].Description, g.getActions()
}

func (g *Game) Act(command string) (msg string, actions []string) {
	parts := strings.Split(command, " ")
	action := parts[0]
	args := parts[1:]

	switch len(args) {
	case 0:
		msg = g.ActZero(action)
	case 1:
		msg = g.ActOne(action, args[0])
	default:
		msg = fmt.Sprintf("Invalid command: %s", command)
	}

	return msg, g.getActions()
}

func (g *Game) ActZero(action string) string {
	if !g.isActiveAction(action) {
		return fmt.Sprintf("Invalid action: %s", action)
	}

	act := g.Actions[action]
	if act.Args != 0 {
		return act.Error
	}
	if act.Type != "location" { // We don't know how to handle this type yet
		return fmt.Sprintf("Type %s not yet supported", act.Type)
	}

	// Act on location
	return g.Locations[g.ActiveLocation].Description
}

func (g *Game) ActOne(action string, arg1 string) string {
	if !g.isActiveAction(action) {
		return fmt.Sprintf("Invalid action: %s", action)
	}

	act := g.Actions[action]
	if act.Args != 1 {
		return act.Error
	}
	if act.Type != "object" { // We don't know how to handle this type yet
		return fmt.Sprintf("Type %s not yet supported", act.Type)
	}

	// Act on object
	object := arg1
	loc := g.Locations[g.ActiveLocation]
	if !loc.isActiveObjet(object) {
		return fmt.Sprintf("No such thing as %s", object)
	}
	result, ok := loc.Objects[object][action]
	if !ok {
		return fmt.Sprintf(act.Default, arg1)
	}

	var msg string
	switch result.Type {
	case "MSG":
		msg = result.Msg
	case "ADD_OBJECT":
		loc.ActiveObjects = append(loc.ActiveObjects, result.Arg)
		msg = result.Msg
	case "ADD_ACTION":
		g.ActiveActions = append(g.ActiveActions, result.Arg)
		msg = result.Msg
	case "SET_LOCATION":
		g.ActiveLocation = result.Arg
		msg = g.Locations[g.ActiveLocation].Description
	case "WIN":
		g.victory = true
		g.ActiveActions = nil
		msg = result.Msg
	}

	return msg
}

func (g *Game) getActions() (actions []string) {
	// Get available 0-arg actions
	for _, act := range g.ActiveActions {
		if g.Actions[act].Args == 0 {
			actions = append(actions, act)
		}
	}

	// Get available 1-arg object actions
	loc := g.Locations[g.ActiveLocation]
	for _, obj := range loc.ActiveObjects {
		for _, act := range g.ActiveActions {
			if g.Actions[act].Args == 1 && g.Actions[act].Type == "object" {
				actions = append(actions, act+" "+obj)
			}
		}
	}
	return actions
}

func (g *Game) isActiveAction(action string) bool {
	for _, act := range g.ActiveActions {
		if act == action {
			return true
		}
	}
	return false
}

type Action struct {
	Args    int    `json:"args"`
	Type    string `json:"type"`
	Default string `json:"default"`
	Error   string `json:"error"`
}

type Location struct {
	Description   string                        `json:"desc"`
	ActiveObjects []string                      `json:"start_objects"`
	Objects       map[string]map[string]*Result `json:"objects"`
}

func (l *Location) isActiveObjet(object string) bool {
	for _, obj := range l.ActiveObjects {
		if obj == object {
			return true
		}
	}
	return false
}

type Result struct {
	Type string `json:"type"`
	Arg  string `json:"arg"`
	Msg  string `json:"msg"`
}
