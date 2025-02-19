package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(conn *websocket.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	for{
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Message received: ", message)
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	}
}

func handleUpgrade(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	log.Println("Client connected")
	handleConnections(conn)
}

func main() {
	conn := "https://myuser:mypassword@magicfights-postgres-1:5432/magic_fights?sslmode=disable"
	db, error := sql.Open("postgres", conn)

	defer db.Close()

	if error != nil {
		log.Fatal(error)
	}
	http.HandleFunc("/ws", handleUpgrade)
	fmt.Println("Server has started on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}