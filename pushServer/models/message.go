package models

// Message Message
type Message struct {
	Title   string
	Body    string
	Badge   int
	Headers map[string]string
}

// ApplyMessage ApplyMessage
func ApplyMessage(title, body string, badge int) *Message {
	return &Message{Title: title, Body: body, Badge: badge, Headers: map[string]string{"apns-priority": "10"}}
}
