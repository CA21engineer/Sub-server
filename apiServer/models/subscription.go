package models

import (
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

type Subscription struct {
	SubscriptionId string
	IconId         string
	Icon           *Icon `gorm:"-"`
	ServiceName    string
	ServiceType    subscription.ServiceType
	Price          int32
	Cycle          int32
	IsOriginal     bool
	FreeTrial      int32
}
