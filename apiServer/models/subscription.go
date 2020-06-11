package models

import (
	"fmt"
	"time"

	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/google/uuid"
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

// NewSubscription new Subscription struct
func NewSubscription(iconID, serviceName string, price, cycle, freeTrial int32) *Subscription {
	uid, _ := uuid.NewUUID()
	return &Subscription{
		SubscriptionID: uid.String(),
		IconID:         iconID,
		ServiceName:    serviceName,
		Price:          price,
		Cycle:          cycle,
		IsOriginal:     false,
		FreeTrial:      freeTrial,
	}
}

// NewSubscriptionToUserSubscription Subscription to new UserSubscription struct
func (s *Subscription) NewSubscriptionToUserSubscription(userID string, startedAt time.Time) *UserSubscription {
	uid, _ := uuid.NewUUID()

	return &UserSubscription{
		UserSubscriptionID: uid.String(),
		UserID:             userID,
		SubscriptionID:     s.SubscriptionID,
		Cycle:              s.Cycle,
		Price:              s.Price,
		StartedAt:          startedAt,
	}
}

// All 登録されている全てのsubscriptionを返す
func (s *Subscription) All() ([]*SubscriptionWithIcon, error) {

	var subscriptionsWithIcon []*SubscriptionWithIcon

	err := DB.Table("subscriptions").
		Select("subscriptions.*, icons.*").
		Joins("join icons on subscriptions.icon_id = icons.icon_id").
		Scan(&subscriptionsWithIcon).
		Error

	if err != nil {
		return nil, err
	}

	return subscriptionsWithIcon, nil
}

func (s *Subscription) PopulerAll() ([]*SubscriptionWithIcon, error) {
	var subscriptionsWithIcon []*SubscriptionWithIcon

	sql := fmt.Sprint(`
		select
			subscriptions.*,
			icons.icon_uri
		from
			user_subscriptions
		inner join
			subscriptions on subscriptions.subscription_id = user_subscriptions.subscription_id
		inner join
				icons on subscriptions.icon_id = icons.icon_id
		where
			subscriptions.is_original = true
		group by
			subscriptions.subscription_id,
			subscriptions.service_name,
			subscriptions.service_type,
			subscriptions.price,
			subscriptions.cycle,
			subscriptions.is_original,
			subscriptions.free_trial,
			icons.icon_uri
		order by
			count(user_subscriptions.subscription_id) DESC
	`)
	if err := DB.Raw(sql).Scan(&subscriptionsWithIcon).Error; err != nil {
		return nil, err
	}

	return subscriptionsWithIcon, nil
}

// Create create original subscription
func (s *Subscription) Create(userID string, startedAt time.Time) error {
	// トランザクション開始
	tx := DB.Begin()

	// ユーザーオリジナルサブスクリプションを作成
	if err := tx.Create(s).Error; err != nil {
		tx.Rollback()
		return err
	}

	// ユーザーオリジナルサブスクリプションをユーザーリストに登録
	userSubscription := s.NewSubscriptionToUserSubscription(userID, startedAt)
	if err := tx.Create(userSubscription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// コミット
	return tx.Commit().Error
}
