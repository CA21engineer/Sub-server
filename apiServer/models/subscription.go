package models

import (
	"time"

	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// Subscription Subscription struct
type Subscription struct {
	SubscriptionID string
	IconID         string
	ServiceName    string
	ServiceType    subscription.ServiceType
	Price          int32
	Cycle          int32
	IsOriginal     bool
	FreeTrial      int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// SubscriptionWithIcon SubscriptionWithIcon struct
type SubscriptionWithIcon struct {
	Subscription
	Icon
}
