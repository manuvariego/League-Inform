package api

import (
	"sync"

	"github.com/gorilla/websocket"
)

//WebSocket structs

type WSInfo struct {
	m    sync.Mutex
	Seq  *int64
	Conn *websocket.Conn
}

type EventPayload struct {
	OpCode    int           `json:"op"`
	Data      HeartBeatData `json:"d"`
	SeqNumber int64         `json:"s"`
	Name      string        `json:"t"`
}

type HeartBeatData struct {
	HeartBeat float64 `json:"heartbeat_interval"`
}

type IdentifyData struct {
	Token   string             `json:"token"`
	Prop    IdentifyProperties `json:"properties"`
	Intents int                `json:"intents"`
}

type IdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

// type ReadyData struct {
// Version int `json:"v"`
// User string
// Guilds string
// SessId string
// ResumeUrl
// }
