package api

import (
	"fmt"
	"os"
)

func (ws *WSInfo) ManageEvent(opCode int) {
	switch opCode {

	case 10:
		// go ws.Heartbeat(ev.Data.HeartBeat)
	default:
		fmt.Println("manageEvent func defaulted, running intent")
		ws.Identify()
		return
	}
}

// func (ws *WSInfo) Write() {
// }

func (ws *WSInfo) Identify() {
	token := os.Getenv("DISCORD_KEY")
	type ident struct {
		Op   int `json:"op"`
		Data IdentifyData
	}
	var data IdentifyData
	data.Token = token
	data.Intents = 7
	data.Prop.Os = "linux"
	data.Prop.Device = "my_lib"
	data.Prop.Browser = "my_lib"

	ws.m.Lock()
	err := ws.Conn.WriteJSON(ident{2, data})
	if err != nil {
		fmt.Println(err)
	}
	ws.m.Unlock()

}
