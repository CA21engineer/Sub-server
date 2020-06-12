package models

import "time"

// NotificationOpt NotificationOpt
type NotificationOpt struct {
	MessageGen func(string) *Message // token => models.Message
	Duration   time.Duration
}

// DefaultNotificationOpt DefaultNotificationOpt
func DefaultNotificationOpt() NotificationOpt {
	return NotificationOpt{
		MessageGen: func(s string) *Message {
			return ApplyMessage("Default Title", "Default Body", 0)
		},
		Duration: time.Second * 10,
	}
}
