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
