package models

import (
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// Subscription Subscription struct
type Subscription struct {
	SubscriptionID string
	IconURI        string
	ServiceName    string
	ServiceType    subscription.ServiceType
	Price          int32
	Cycle          int32
	IsOriginal     bool
	FreeTrial      int32
}

func (s *Subscription) All() ([]*Subscription, error) {
	var subscriptions []*Subscription

	err := DB.Table("subscriptions").
		Select("subscriptions.subscription_id, icons.icon_uri, subscriptions.service_name, subscriptions.service_type, " +
			"subscriptions.service_type, subscriptions.price, subscriptions.cycle, subscriptions.is_original, subscriptions.free_trial").
		Joins("left outer join icons on subscriptions.icon_id = icons.icon_id").
		Scan(&subscriptions).
		Error

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}
