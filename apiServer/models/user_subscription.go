package models

import "time"

type UserSubscription struct {
	UserSubscriptionId string
	UserId             string
	Icon               Icon `gorm:"-"`
	SubscriptionId     string
	Subscription       Subscription `gorm:"-"`
	Cycle              int32
	Price              int32
	StartedAt          time.Time
}

func (u *UserSubscription) GetUserSubscriptions(userId string) ([]*UserSubscription, error) {
	var userSubscriptions []*UserSubscription
	if err := DB.Where("user_id = ?", userId).Find(&userSubscriptions).Error; err != nil {
		return nil, err
	}

	for i, v := range userSubscriptions {
		var icon Icon
		var subscription Subscription
		DB.Where("subscription_id = ?", v.SubscriptionId).Find(&subscription)
		DB.Where("icon_id = ?", subscription.IconId).Find(&icon)
		userSubscriptions[i].Subscription = subscription
		userSubscriptions[i].Icon = icon
	}
	return userSubscriptions, nil
}
