package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func eventsHandler(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	challenge := gjson.Get(string(body), "challenge").String()
	c.String(http.StatusOK, challenge)
}
