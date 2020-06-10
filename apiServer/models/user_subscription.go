package models

import "time"

type UserSubscription struct {
	UserSubscriptionId string
	UserId             string
	SubscriptionId     string
	Cycle              int32
	Price              int32
	StartedAt          time.Time
}
