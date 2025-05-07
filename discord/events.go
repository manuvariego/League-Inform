package discord

import (
	"encoding/json"
	"fmt"
)

func (ws *WSInfo) ManageEvent(eventName string, payload []byte) {
	switch eventName {

	case "MESSAGE_CREATE":
		var msgcreate Message

		fmt.Println(string(payload))
		err := json.Unmarshal(payload, &msgcreate)
		if err != nil {
			fmt.Println("Error while decoding", err)
		}
		HandleMessage(msgcreate.Content)

		// fmt.Println("--")
		// fmt.Println(msgcreate)
		// fmt.Println("--")
	}

}

func (ws *WSInfo) ManageOpCode(opCode int, eventName string, data []byte) {
	switch opCode {

	case OpCodeHeartbeat:
		hbData := &HeartBeatData{}

		err := json.Unmarshal(data, &hbData)
		if err != nil {
			fmt.Println(err)
			return
		}

		go ws.Heartbeat(hbData.HeartBeat)
		ws.Identify()

	case OpCodeTypeTwo:

	case OpCodeManageEvent:
		ws.ManageEvent(eventName, data)

	default:
		fmt.Println("manageEvent func defaulted, running intent")
		return
	}
}
