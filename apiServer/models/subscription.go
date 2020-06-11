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

// All 登録されている全てのsubscriptionを返す
func (s *Subscription) All() ([]*SubscriptionWithIcon, error) {

	var subscriptionsWithIcon []*SubscriptionWithIcon

	err := DB.Table("subscriptions").
		//Select("subscriptions.subscription_id, icons.icon_uri, subscriptions.service_name, subscriptions.service_type, " +
		//	"subscriptions.service_type, subscriptions.price, subscriptions.cycle, subscriptions.is_original, subscriptions.free_trial").
		Select("subscriptions.*, icons.*").
		Joins("left outer join icons on subscriptions.icon_id = icons.icon_id").
		Scan(&subscriptionsWithIcon).
		Error

	if err != nil {
		return nil, err
	}

	return subscriptionsWithIcon, nil
}
