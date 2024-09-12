package api

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"

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
	//initialWaitTime in float, <<rand float between 0-1>>
	initialTime := heartbeat * (rand.Float64())
	initialDelay := time.Duration(initialTime) * time.Millisecond
	count := 0

	type heartbeats struct {
		Op  int   `json:"op"`
		Seq int64 `json:"d"`
	}

	if count < 1 {
		//Ticker for the first heartbeat
		ticker := time.NewTicker(initialDelay)
		fmt.Println("Inside first heartbeat")
		defer ticker.Stop()
		select {
		case <-ticker.C:
			count += 1
			seqNumber := atomic.LoadInt64(ws.Seq)
			ws.m.Lock()
			ws.Conn.WriteJSON(heartbeats{1, seqNumber})
			ws.m.Unlock()
		}
	}
	//Ticker for the heartbeats not including the first one
	heartbeatInterval := time.Duration(heartbeat) * time.Millisecond
	ticker := time.NewTicker(heartbeatInterval)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Inside constant heartbeat")
			seqNumber := atomic.LoadInt64(ws.Seq)
			ws.m.Lock()
			ws.Conn.WriteJSON(heartbeats{1, seqNumber})
			ws.m.Unlock()
		}

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

		atomic.StoreInt64(ws.Seq, ev.SeqNumber)
		ws.ManageEvent(ev)
	}
}

// func Writer(conn *websocket.Conn, mess Message) {
// 	w, err := conn.NextWriter
// }
