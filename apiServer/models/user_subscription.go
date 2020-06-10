package models

import (
	"fmt"
	"time"

	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

//UserSubscription struct
type UserSubscription struct {
	UserSubscriptionID string
	SubscriptionID     string
	ServiceType        subscription.ServiceType
	IconURI            string
	ServiceName        string
	Price              int32
	Cycle              int32
	FreeTrial          int32
	IsOriginal         bool
	StartedAt          time.Time
}

// GetUserSubscriptions 特定ユーザーの登録しているsubscriptionを返す
func (u *UserSubscription) GetUserSubscriptions(userID string) ([]*UserSubscription, error) {
	var userSubscriptions []*UserSubscription
	sql := fmt.Sprintf("select user_subscriptions.user_subscription_id,user_subscriptions.user_id,icons.*,user_subscriptions.subscription_id,subscriptions.*,user_subscriptions.cycle,user_subscriptions.price from user_subscriptions join subscriptions on user_subscriptions.subscription_id = subscriptions.subscription_id left outer join icons on subscriptions.icon_id = icons.icon_id where user_subscriptions.user_id = '%s'", userID)
	if err := DB.Raw(sql).Scan(&userSubscriptions).Error; err != nil {
		return nil, err
	}
	return userSubscriptions, nil
}
