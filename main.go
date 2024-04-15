package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var lightHouseIntervalSeconds = "1"
var wsMessage = "nolight"

func main() {
	router := gin.Default()
	router.GET("/ping", Ping)
	router.GET("/interval", GetInterval)
	router.GET("/ws", StartWebSocket)
	router.Run(":3000")
}

func Timer() {
	for {
		wsMessage = "light"
		time.Sleep(time.Second)
		wsMessage = "nolight"
	}
}

func StartWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("light"))
		time.Sleep(time.Second / 2)
		conn.WriteMessage(websocket.TextMessage, []byte("lights"))
		time.Sleep(time.Second / 2)
	}
}

func Ping(c *gin.Context) {

	lightHouseIntervalSeconds = c.Query("lightHouseTime")
	c.String(http.StatusOK, lightHouseIntervalSeconds)
}

func GetInterval(c *gin.Context) {
	c.String(http.StatusOK, lightHouseIntervalSeconds)
}
