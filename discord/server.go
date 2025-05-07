package discord

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
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

func (ws *WSInfo) Identify() {
	token := os.Getenv("DISCORD_KEY")
	fmt.Println(token)

	if token == "" {
		fmt.Println("token null")
	}

	type ident struct {
		Op   int          `json:"op"`
		Data IdentifyData `json:"d"`
	}

	var data IdentifyData
	data.Token = token
	data.Intents = 33281
	data.Prop.Os = "linux"
	data.Prop.Device = "my_lib"
	data.Prop.Browser = "my_lib"
	fmt.Println(data)

	ws.m.Lock()
	err := ws.Conn.WriteJSON(ident{2, data})
	if err != nil {
		fmt.Println("test")
		fmt.Println(err)
	}
	ws.m.Unlock()
}

func (ws *WSInfo) Reader() {

	for {

		var payload DefaultPayload

		_, p, err := ws.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		err = json.Unmarshal(p, &payload)
		if err != nil {
			fmt.Println(err)
		}

		//Temp : Prints payload
		// fmt.Println(payload)

		atomic.StoreInt64(ws.Seq, payload.SeqNumber)

		ws.ManageOpCode(payload.Opcode, payload.Name, payload.Data)
	}
}
