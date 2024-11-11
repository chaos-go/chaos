package event_gateway

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

const DiscordBaseUrl = "wss://gateway.discord.gg"
const DiscordApiVersion = "10"
const DiscordEncoding = "json"

type EventGateway struct {
	ws               *websocket.Conn
	connectionString string
}

func createConnectionString(baseUrl string, discordVersion string, encoding string, compress bool) string {
	v := url.Values{}
	v.Set("v", discordVersion)
	v.Set("encoding", encoding)
	v.Set("compress", fmt.Sprint(compress))
	v.Encode()
	connectionString, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	connectionString.RawQuery = v.Encode()
	return connectionString.String()
}

// New returns new EventGateway structure for communicating with Discord Event Gateway
// (https://discord.com/developers/docs/events/gateway)
func New() *EventGateway {
	return &EventGateway{
		connectionString: createConnectionString(DiscordBaseUrl, DiscordApiVersion, DiscordEncoding, true),
	}
}

const (
	TextMessage     = 1
	BinaryMessage   = 2
	PingPongMessage = 3
)

type MessageType = int

// TODO: rewrite for sending events as structs
func (g *EventGateway) sendEvent(op int, d any, s int, t string, msgType MessageType) error {
	payload := &EventPayload{Op: op, D: d, S: s, T: t}
	// 1, text message
	//2, binary message
	//3, Ping message and Pong message
	encodedMsg, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = g.ws.WriteMessage(msgType, encodedMsg)
	if err != nil {
		return err
	}
	return nil
}
