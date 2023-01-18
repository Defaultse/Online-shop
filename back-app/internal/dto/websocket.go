package dto

type WebsocketMsg struct {
	Type    float64     `json:"type"`
	Message interface{} `json:"message"`
}

type WebsocketMsg2 struct {
	Type float64     `json:"type"`
	Body interface{} `json:"body"`
}
