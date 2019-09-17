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
