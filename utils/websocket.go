package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(conn *websocket.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	for {
		_, messageBytes, err := conn.ReadMessage()
		message := string(messageBytes)

		if err != nil {
			fmt.Printf("30: %s", err)
			break
		}

		// res := router(message)
		fmt.Println("Message received: ", message)
		if err := conn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
			fmt.Printf("38: %s", err)
			break
		}

	}
}

func HandleUpgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	log.Println("Client connected")
	HandleConnections(conn)
}
