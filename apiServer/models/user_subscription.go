package models

import "time"

type UserSubscription struct {
	UserSubscriptionId string
	UserId             string
	Subscription       *Subscription
	Cycle              int32
	Price              int32
	StartedAt          time.Time
}
