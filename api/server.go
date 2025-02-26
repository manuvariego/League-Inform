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

func (ws *WSInfo) Write() {
	s := "hey brother"
	bytesArray := []byte(s)
	err := ws.Conn.WriteMessage(1, bytesArray)
	if err != nil {
		fmt.Println(err)
	}

}

func (ws *WSInfo) Reader() string {

	for {
		type eventPayload struct {
			Opcode    int    `json:"op"`
			SeqNumber int64  `json:"s"`
			Name      string `json:"t"`
		}

		var event eventPayload

		_, p, err := ws.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(p, &event)
		if err != nil {
			fmt.Println(err)
		}

		//Temp : Prints event payload
		fmt.Println(event)

		atomic.StoreInt64(ws.Seq, event.SeqNumber)

		//Temp : sends event payload to manageEvent
		ws.ManageEvent(event.Opcode, p)
	}
}
