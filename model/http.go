package model

const (
	ContentTypeJSON = "application/json"
)

type ResponsePayload struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}

type GetStartPayload struct {
	Payload string `json:"payload"`
}
type RequestBodyCreateGetStart struct {
	GetStart GetStartPayload `json:"get_started"`
}

type GreetingMessage struct {
	Locate string `json:"locale"`
	Text   string `json:"text"`
}
type RequestBodyCreateGreetingMessage struct {
	GreetingMessages []GreetingMessage `json:"greeting"`
}

type CallToAction struct {
	Type               string  `json:"type"`
	Title              *string `json:"title,omitempty"`
	Payload            *string `json:"payload,omitempty"`
	URL                *string `json:"url,omitempty"`
	WebviewHeightRatio *string `json:"webview_height_ratio,omitempty"`
}
type PersistentMenu struct {
	Locale                string         `json:"locale"`
	ComposerInputDisabled bool           `json:"composer_input_disabled"`
	CallToActions         []CallToAction `json:"call_to_actions"`
}
type RequestBodyCreatePersistentMenu struct {
	PSID            string           `json:"psid"`
	PersistentMenus []PersistentMenu `json:"persistent_menu"`
}

type GenericElement struct {
	Title         string         `json:"title"`
	ImageUrl      string         `json:"image_url"`
	Subtitle      string         `json:"subtitle"`
	DefaultAction CallToAction   `json:"default_action"`
	Buttons       []CallToAction `json:"buttons"`
}
type TemplateGenericPayload struct {
	TemplateType string           `json:"template_type"`
	Elements     []GenericElement `json:"elements"`
}
type TemplateGenericAttachment struct {
	Type    string                 `json:"type"`
	Payload TemplateGenericPayload `json:"payload"`
}
type TemplateGenericMessage struct {
	Attachment TemplateGenericAttachment `json:"attachment"`
}
type RequestBodyCreateGenericTemplate struct {
	Recipient Recipient              `json:"recipient"`
	Message   TemplateGenericMessage `json:"message"`
}

type ReceiptElement struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Currency string `json:"currency"`
	ImageUrl string `json:"image_url"`
}
type ReceiptAdjustments struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
type ReceiptSummary struct {
	Subtotal     float64 `json:"subtotal"`
	ShippingCost float64 `json:"shipping_cost"`
	TotalTax     float64 `json:"total_tax"`
	TotalCost    float64 `json:"total_cost"`
}
type ReceiptAddress struct {
	Street     string `json:"street_1"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Country    string `json:"country"`
}
type TemplateReceiptPayload struct {
	TemplateType  string                `json:"template_type"`
	RecipientName string                `json:"recipient_name"`
	OrderName     string                `json:"order_number"`
	Currency      string                `json:"currency"`
	PaymentMethod string                `json:"payment_method"`
	OrderUrl      string                `json:"order_url"`
	Timestamp     string                `json:"timestamp"`
	Address       *ReceiptAddress       `json:"address,omitempty"`
	Summary       ReceiptSummary        `json:"summary"`
	Adjustments   *[]ReceiptAdjustments `json:"adjustments,omitempty"`
	Elements      []ReceiptElement      `json:"elements"`
}
type TemplateReceiptAttachment struct {
	Type    string                 `json:"type"`
	Payload TemplateReceiptPayload `json:"payload"`
}
type TemplateReceiptMessage struct {
	Attachment TemplateReceiptAttachment `json:"attachment"`
}
type RequestBodyCreateReceiptTemplate struct {
	Recipient Recipient              `json:"recipient"`
	Message   TemplateReceiptMessage `json:"message"`
}

type TemplateButtonPayload struct {
	TemplateType string         `json:"template_type"`
	Text         string         `json:"text"`
	Button       []CallToAction `json:"buttons"`
}
type TemplateButtonAttachment struct {
	Type    string                `json:"type"`
	Payload TemplateButtonPayload `json:"payload"`
}
type TemplateButtonMessage struct {
	Attachment TemplateButtonAttachment `json:"attachment"`
}
type RequestBodyCreateButtonTemplate struct {
	Recipient Recipient             `json:"recipient"`
	Message   TemplateButtonMessage `json:"message"`
}

type QuickReply struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Payload     string `json:"payload"`
}
type QuickReplyMessage struct {
	Text         string       `json:"text"`
	QuickReplies []QuickReply `json:"quick_replies"`
}
type RequestBodyCreateQuickReply struct {
	Recipient   Recipient         `json:"recipient"`
	MessageType string            `json:"messaging_type"`
	Message     QuickReplyMessage `json:"message"`
}
