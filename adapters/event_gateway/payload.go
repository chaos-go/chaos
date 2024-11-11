package event_gateway

type EventPayload struct {
	Op int `json:"op"`
	// d could be any json value
	D any    `json:"d"`
	S int    `json:"s,omitempty"`
	T string `json:"t,omitempty"`
}
