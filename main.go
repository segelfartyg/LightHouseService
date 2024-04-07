package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var lightHouseIntervalSeconds = "1"

func main() {
	router := gin.Default()
	router.GET("/ping", Ping)

	router.GET("/interval", GetInterval)

	router.Run("localhost:3000")
}

func Ping(c *gin.Context) {

	lightHouseIntervalSeconds = c.Query("lightHouseTime")
	c.String(http.StatusOK, lightHouseIntervalSeconds)
}

func GetInterval(c *gin.Context) {
	c.String(http.StatusOK, lightHouseIntervalSeconds)
}
