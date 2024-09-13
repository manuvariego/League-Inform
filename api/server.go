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

	return ws
}

func (ws *WSInfo) Heartbeat(heartbeat float64) {

	firstHB := true

	type heartbeats struct {
		Op  int   `json:"op"`
		Seq int64 `json:"d"`
	}

	if firstHB {

		fmt.Println("Inside first heartbeat")
		//Timer for the first heartbeat
		initialDelay := time.Duration(heartbeat*(rand.Float64())) * time.Millisecond
		time.Sleep(initialDelay)

		seqNumber := atomic.LoadInt64(ws.Seq)
		ws.m.Lock()
		err := ws.Conn.WriteJSON(heartbeats{1, seqNumber})
		if err != nil {
			fmt.Println(err)
		}
		ws.m.Unlock()
	}

	//Ticker for the heartbeats not including the first one
	heartbeatInterval := time.Duration(heartbeat) * time.Millisecond
	ticker := time.NewTicker(heartbeatInterval)

	for range ticker.C {
		fmt.Println("Inside constant heartbeat")
		seqNumber := atomic.LoadInt64(ws.Seq)
		ws.m.Lock()
		err := ws.Conn.WriteJSON(heartbeats{1, seqNumber})
		if err != nil {
			fmt.Println(err)
		}
		ws.m.Unlock()
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

		//Temp : sends event payload to manageEvent
		ws.ManageEvent(ev)
	}
}

// func Writer(conn *websocket.Conn, mess Message) {
// 	w, err := conn.NextWriter
// }
