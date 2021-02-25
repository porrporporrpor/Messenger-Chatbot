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

type Element struct {
	Title         string         `json:"title"`
	ImageUrl      string         `json:"image_url"`
	Subtitle      string         `json:"subtitle"`
	DefaultAction CallToAction   `json:"default_action"`
	Buttons       []CallToAction `json:"buttons"`
}
type TemplateAttachmentPayload struct {
	TemplateType string    `json:"template_type"`
	Elements     []Element `json:"elements"`
}
type TemplateAttachment struct {
	Type    string                    `json:"type"`
	Payload TemplateAttachmentPayload `json:"payload"`
}
type TemplateMessage struct {
	Attachment TemplateAttachment `json:"attachment"`
}
type RequestBodyCreateGenericTemplate struct {
	Recipient Recipient       `json:"recipient"`
	Message   TemplateMessage `json:"message"`
}
