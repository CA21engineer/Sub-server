package models

import (
	"fmt"
	"time"
)

//UserSubscription struct
type UserSubscription struct {
	UserSubscriptionID string
	UserID             string
	SubscriptionID     string
	Price              int32
	Cycle              int32
	StartedAt          time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Subscription       Subscription
	Icon               Icon
}

// UserSubscriptionDiff UserSubscriptionDiff struct
type UserSubscriptionDiff struct {
	UserSubscriptionID string
	UserID             string
	StartedAt          time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// Record Record struct
type Record struct {
	UserSubscriptionDiff
	Subscription
	Icon
}

// GetUserSubscriptions 特定ユーザーの登録しているsubscriptionを返す
func (u *UserSubscription) GetUserSubscriptions(userID string) ([]*UserSubscription, error) {

	var records []*Record
	sql := fmt.Sprintf("select user_subscriptions.user_subscription_id,user_subscriptions.user_id,user_subscriptions.price, user_subscriptions.cycle, user_subscriptions.started_at,user_subscriptions.created_at,user_subscriptions.updated_at,subscriptions.*,icons.* from user_subscriptions join subscriptions on user_subscriptions.subscription_id = subscriptions.subscription_id join icons on subscriptions.icon_id = icons.icon_id where user_subscriptions.user_id = '%s' ORDER BY user_subscriptions.updated_at asc", userID)
	if err := DB.Raw(sql).Scan(&records).Error; err != nil {
		return nil, err
	}

	userSubscriptions := make([]*UserSubscription, len(records))
	for i, e := range records {
		userSubscriptions[i] = &UserSubscription{
			UserSubscriptionID: e.UserSubscriptionDiff.UserSubscriptionID,
			UserID:             e.UserSubscriptionDiff.UserID,
			Price:              e.Price,
			Cycle:              e.Cycle,
			StartedAt:          e.UserSubscriptionDiff.StartedAt,
			CreatedAt:          e.UserSubscriptionDiff.CreatedAt,
			UpdatedAt:          e.UserSubscriptionDiff.UpdatedAt,
			Subscription:       e.Subscription,
			Icon:               e.Icon,
		}
	}
	return userSubscriptions, nil

}

// Find 特定のuser_subscriptionを返す
func (u *UserSubscription) Find(userSubscriptionID string) (*UserSubscription, error) {
	var userSubscription UserSubscription
	if err := DB.Where("user_subscription_id = ?", userSubscriptionID).Find(&userSubscription).Error; err != nil {
		return nil, err
	}
	return &userSubscription, nil
}
