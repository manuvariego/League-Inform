package api

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// uri := "wss://gateway.discord.gg"

//	var upgrader = websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//		CheckOrigin: func(r *http.Request) bool {
//			origin := r.Header.Get("Origin")
//			return origin == "gateway.discord.gg"
//
//		},
//	}

type Disc struct {
	HeartBeat int `json:"heartbeat_interval"`
}

type Message struct {
	Op     int  `json:"op"`
	Discod Disc `json:"d"`
}

func ConnectToDiscord() *websocket.Conn {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	if err != nil {
		fmt.Println(err)
	}
	// messageType, p, err := conn.NextRear()
	// fmt.Println(messageType)
	// fmt.Println(string(p))
	//
	return conn
}

func Reader(conn *websocket.Conn) string {
	for {
		m := Message{}
		messageType, r, err := conn.NextReader()
		json.NewDecoder(r).Decode(&m)
		fmt.Println(messageType)
		fmt.Println(m)
		fmt.Println(err)
	}
}
