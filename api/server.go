package api

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func ConnectToDiscord() *WSInfo {
	ws := NewWSConnection()
	//Returns a discord websocket connection
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	if err != nil {
		fmt.Println(err)
	}
	ws.Conn = conn

	// go Reader(conn)
	// // go Writer(conn)

	return ws
}

func (ws *WSInfo) Heartbeat(heartbeat float64) {
	// initialTime := heartbeat * (rand.Float64())
	for {

	}

}

func (ws *WSInfo) Reader(conn *websocket.Conn) string {

	for {
		m := EventPayload{}

		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(p, &m)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(m)
		fmt.Println(m.SeqNumber)

		// json.NewDecoder(bytes.NewReader(p)).Decode(&m)

		// ws.seq = m.SeqNumber

		// op := m.OpCode
		// seq := m.SeqNumber
		// switch op {
		// case 10:
		// 	go Heartbeat(conn, m, seq)
		// }

		// go Writer(conn, m)

		// fmt.Println(m)
		// fmt.Println(err)
	}
}

// func Writer(conn *websocket.Conn, mess Message) {
// 	w, err := conn.NextWriter
// }
