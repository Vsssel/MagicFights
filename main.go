package main

import (
	"MagicFights/utils"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	utils.ConnectDB()
	http.HandleFunc("/ws", utils.HandleUpgrade)
	fmt.Println("Server has started on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
