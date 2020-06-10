package models

import (
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

const (
	NOT_FOUND = iota + 1
	NOT_CATEGORIZED
	MUSIC
	MOVIE
	MATCHING
	STORAGE
)

type Subscription struct {
	SubscriptionId string
	Icon           *Icon
	ServiceName    string
	ServiceType    subscription.ServiceType
	Price          int32
	Cycle          int32
	IsOriginal     bool
	FreeTrial      int32
}
