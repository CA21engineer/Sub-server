package models

import (
	"context"
	"time"
)

// NotificationCrawler NotificationCrawler
type NotificationCrawler struct {
	PushNotification *PushNotification
	Option           NotificationCrawlerOpt
	Execute          func(context.Context, *PushNotification)
}

// StartCrawlerTimer StartCrawlerTimer
func (n NotificationCrawler) StartCrawlerTimer(ctx context.Context) {
	ticker := time.NewTicker(n.Option.Duration)
	defer ticker.Stop()
	n.Execute(ctx, n.PushNotification)
	for {
		select {
		case <-ticker.C:
			n.Execute(ctx, n.PushNotification)
		}
	}
}
