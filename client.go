package chaos

import (
	intents "github.com/golang-chaos/chaos/resources"
	"sync"
)

type Identify struct {
	Token string `json:"token"`
	//Properties     IdentifyProperties  `json:"properties"`
	Compress       bool    `json:"compress"`
	LargeThreshold int     `json:"large_threshold"`
	Shard          *[2]int `json:"shard,omitempty"`
	//Presence       GatewayStatusUpdate `json:"presence,omitempty"`
	Intents intents.Intent `json:"intents"`
}

type Client struct {
	// We use composition with RWMutex because ???
	sync.RWMutex

	// Token obtained from Discord Developer Portal
	// (https://discord.com/developers/applications/?ysclid=m3965p4b8p599291877)
	// right now only Bot tokens are supported
	Token string

	// Identify is a payload sent during the Discord gateway handshake
	// (https://discord.com/developers/docs/events/gateway#identifying)
	Identify any //TODO

}
