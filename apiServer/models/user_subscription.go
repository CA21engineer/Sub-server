package models

import "time"

//UserSubscription struct
type UserSubscription struct {
	UserSubscriptionID string
	UserID             string
	Icon               Icon `gorm:"-"`
	SubscriptionID     string
	Subscription       Subscription `gorm:"-"`
	Cycle              int32
	Price              int32
	StartedAt          time.Time
}

// GetUserSubscriptions 特定ユーザーの登録しているsubscriptionを返す
func (u *UserSubscription) GetUserSubscriptions(userID string) ([]*UserSubscription, error) {
	var userSubscriptions []*UserSubscription
	if err := DB.Where("user_id = ?", userID).Find(&userSubscriptions).Error; err != nil {
		return nil, err
	}

	for i, v := range userSubscriptions {
		var icon Icon
		var subscription Subscription
		DB.Where("subscription_id = ?", v.SubscriptionID).Find(&subscription)
		DB.Where("icon_id = ?", icon.IconID).Find(&icon)
		userSubscriptions[i].Subscription = subscription
		userSubscriptions[i].Icon = icon
	}
	return userSubscriptions, nil
}
