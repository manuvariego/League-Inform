package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// uri := "wss://gateway.discord.gg"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "gateway.discord.gg"

	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}

}

func Server() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server started on port 3000 ")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Errorf("Error starting server: %v", err)
	}

}
