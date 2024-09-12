package api

import "fmt"

func (ws *WSInfo) ManageEvent(ev *EventPayload) {
	switch ev.OpCode {

	case 10:
		go ws.Heartbeat(ev.Data.HeartBeat)
	default:
		fmt.Println("manageEvent func defaulted")
		return
	}

}
