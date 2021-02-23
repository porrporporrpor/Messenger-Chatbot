package model

type Sender struct {
	ID string `json:"id"`
}
type Recipient struct {
	ID string `json:"id"`
}
type Message struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp float64   `json:"timestamp"`
}

type Entry struct {
	ID        string  `json:"id"`
	Time      float64 `json:"time"`
	Messaging []Message
}

type MessengerRequestBody struct {
	Entry  []Entry `json:"entry"`
	Object string  `json:"object"`
}
