package model

const (
	TextOperationGetStart = "Get Start!"
)

/**
Media in Message struct can be
- text string : for normal text message and emoji
- attachments interface : for audio, video and image (sticker, gif also seen as image)
*/

/**
audio
type : audio
payload : url

video
type : video
payload : url


gif
type : image
payload : url

image
type : image
payload : url

sticker
type : image
payload : sticker_id , url
*/

type Payload struct {
	StickerID *string `json:"sticker_id,omitempty"`
	URL       string  `json:"url"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Sender struct {
	ID string `json:"id"`
}
type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	MID         string        `json:"mid,omitempty"`
	Text        string        `json:"text,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
	QuickReply  *struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
}

type Postback struct {
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp float64   `json:"timestamp"`
	Message   *Message  `json:"message,omitempty"`
	Postback  *Postback `json:"postback,omitempty"`
}

type Entry struct {
	ID        string       `json:"id"`
	Time      float64      `json:"time"`
	Messaging *[]Messaging `json:"messaging"`
}

type MessengerRequestBody struct {
	Entry  []Entry `json:"entry"`
	Object string  `json:"object"`
}
