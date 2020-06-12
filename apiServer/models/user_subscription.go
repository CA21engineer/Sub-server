package models

import (
	"fmt"
	"github.com/google/uuid"
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

// NewUserSubscription new UserSubscription struct
func NewUserSubscription(userID, subscriptionID string, price, cycle int32, startedAt time.Time) *UserSubscription {
	return &UserSubscription{
		UserID:         userID,
		SubscriptionID:    subscriptionID,
		Price:          price,
		Cycle:          cycle,
		StartedAt: startedAt,
	}
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

//Register user_subscriptionを新規作成する
func (u *UserSubscription) Register() error {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u.UserSubscriptionID = uid.String()
	if err = DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// Unregister 特定のuser_subscriptionを削除する
func (u *UserSubscription) Unregister(userID string, userSubscriptionID string) (*UserSubscription, error) {
	var userSubscription UserSubscription
	findUserSubscriptionQuery := DB.Where("user_id = ? and user_subscription_id = ?", userID, userSubscriptionID)
	if err := findUserSubscriptionQuery.First(&userSubscription).Error; err != nil {
		return nil, err
	}
	findUserSubscriptionQuery.Delete(&userSubscription)
	return &userSubscription, nil
}
