package discord

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	OpCodeHeartbeat   = 10
	OpCodeTypeTwo     = 2
	OpCodeManageEvent = 0
)

//WebSocket connection wrapper

type WSInfo struct {
	m    sync.Mutex
	Seq  *int64
	Conn *websocket.Conn
}

// Default discord payload structure
type DefaultPayload struct {
	Opcode    int             `json:"op"`
	SeqNumber int64           `json:"s"`
	Name      string          `json:"t"`
	Data      json.RawMessage `json:"d"`
}

// Discord Data structures
type HeartBeatData struct {
	HeartBeat float64 `json:"heartbeat_interval"`
}

type Message struct {
	Id      string `json:"id"`
	GuildId string `json:"guild_id,omitempty"`
	Author  User   `json:"author"`
	Content string `json:"content"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
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

// ReadyData without the user/guilds/app objects. Add those eventually
type ReadyData struct {
	Version int `json:"v"`
	// User string
	// Guilds string
	SessId    string `json:"session_id"`
	ResumeUrl string `json:"resume_gateway_url"`
	// AppObject string
}
