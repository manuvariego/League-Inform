package api

import (
	"bytes"
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

type HeartBeatData struct {
	HeartBeat float64 `json:"heartbeat_interval"`
}

type EventPayload struct {
	OpCode    int           `json:"op"`
	Data      HeartBeatData `json:"d"`
	SeqNumber int           `json:"s"`
	Name      string        `json:"t"`
}

func ConnectToDiscord() *websocket.Conn {
	//Returns a discord websocket connection
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	if err != nil {
		fmt.Println(err)
	}

	go Reader(conn)
	// go Writer(conn)

	return conn
}

func Heartbeat(conn *websocket.Conn, payload EventPayload, seq int) {
	// heartbeat := payload.Data.HeartBeat
	// initialTime := heartbeat * (rand.Float64())
	for {

	}

}

func Reader(conn *websocket.Conn) string {
	for {
		m := EventPayload{}
		_, p, err := conn.ReadMessage()
		// messageType, r, err := conn.NextReader()
		json.NewDecoder(bytes.NewReader(p)).Decode(&m)
		op := m.OpCode
		seq := m.SeqNumber
		switch op {
		case 10:
			go Heartbeat(conn, m, seq)
		}

		// go Writer(conn, m)

		fmt.Println(m)
		fmt.Println(err)
	}
}

// func Writer(conn *websocket.Conn, mess Message) {
// 	w, err := conn.NextWriter
// }
