package api

import (
	"encoding/json"
	"fmt"
	"os"
	"sync/atomic"
)

func (ws *WSInfo) ManageEvent(opCode int, data []byte) {
	switch opCode {

	case 10:
		hbEvent := &EventPayload{}

		err := json.Unmarshal(data, &hbEvent)
		if err != nil {
			fmt.Println(err)
		}

		atomic.StoreInt64(ws.Seq, hbEvent.SeqNumber)

		go ws.Heartbeat(hbEvent.Data.HeartBeat)
		ws.Identify()

	case 2:
		fmt.Println("test")

	case 0:
		fmt.Println("test")

	default:
		fmt.Println("manageEvent func defaulted, running intent")
		return
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
	data.Intents = 7
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
