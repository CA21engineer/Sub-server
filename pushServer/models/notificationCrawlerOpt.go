package models

import (
	"time"
)

// NotificationCrawlerOpt NotificationCrawlerOpt
type NotificationCrawlerOpt struct {
	MessageGen func(string) *Message // token => models.Message
	Duration   time.Duration
}

// DefaultNotificationCrawlerOpt DefaultNotificationCrawlerOpt
func DefaultNotificationCrawlerOpt() NotificationCrawlerOpt {
	return NotificationCrawlerOpt{
		MessageGen: func(s string) *Message {
			return ApplyMessage("Default Title", "Default Body", 0)
		},
		Duration: time.Second * 10,
	}
}
