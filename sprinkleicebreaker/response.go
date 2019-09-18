package main

import "github.com/gin-gonic/gin"

func startIcebreakerButton() []gin.H {
	return []gin.H{
		button("Start Icebreaker", "start"),
	}
}
func button(text, value string) gin.H {
	return gin.H{
		"name":  value,
		"text":  text,
		"type":  "button",
		"value": value,
	}
}

func iceBreakerSelectMenu() []gin.H {
	return []gin.H{
		gin.H{
			"name": "icebreaker_options",
			"text": "Here are your options...",
			"type": "select",
			"options": []gin.H{
				option("Option 1", "option-1"),
				option("Option 2", "option-2"),
				option("Option 3", "option-3"),
			},
		},
	}
}
func option(text, value string) gin.H {
	return gin.H{
		"text":  text,
		"value": value,
	}
}
