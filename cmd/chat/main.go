package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.GET("/", homePage)
	router.GET("/room/:id", homePage)
	router.GET("/ws", wsEndpoint)

	router.Run(":9999")
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))
		fmt.Println(messageType)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage("id."+string(c.Param("id")), []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}