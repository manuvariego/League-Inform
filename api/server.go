package api

import (
	"encoding/json"
	"fmt"
	"internal/runtime/atomic"
	"math/rand/v2"

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
	initialTime := heartbeat * (rand.Float64())
	count := 0

	type heartbeats struct {
		Op  int   `json:"op"`
		Seq int64 `json:"d"`
	}

	if count < 1 {
		count += 1
		seqNumber := atomic.Loadint64(ws.Seq)
		ws.m.Lock()
		ws.Conn.WriteJSON(heartbeats{1, seqNumber})
		ws.m.Unlock()
	}

	for {

	}

}

func (ws *WSInfo) Reader(conn *websocket.Conn) string {

	for {
		ev := &EventPayload{}

		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(p, &ev)
		if err != nil {
			fmt.Println(err)
		}

		//Temp : Prints event payload
		fmt.Println(ev)

		atomic.Storeint64(ws.Seq, ev.SeqNumber)
		ws.ManageEvent(ev)

	}
}

// func Writer(conn *websocket.Conn, mess Message) {
// 	w, err := conn.NextWriter
// }
