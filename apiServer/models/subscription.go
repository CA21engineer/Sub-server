package models

import (
	"time"

	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/google/uuid"
)

// Subscription Subscription struct
type Subscription struct {
	SubscriptionID string
	IconID         string
	Icon           *Icon `gorm:"-"`
	ServiceName    string
	ServiceType    subscription.ServiceType
	Price          int32
	Cycle          int32
	IsOriginal     bool
	FreeTrial      int32
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
		IsOriginal:     true,
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
