package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// ConvertGRPCSubscriptionResponse `*models.Subscription`を`*subscription.Subscription`に変換
func ConvertGRPCSubscriptionResponse(s *models.Subscription) *subscription.Subscription {
	return &subscription.Subscription{
		SubscriptionId: s.SubscriptionID,
		ServiceType:    s.ServiceType,
		IconUri:        s.IconURI,
		ServiceName:    s.ServiceName,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      s.FreeTrial,
		IsOriginal:     s.IsOriginal,
		StartedAt:      nil,
	}
}

// ConvertGRPCSubscriptionListResponse `[]*models.Subscription`を`[]*subscription.Subscription`に変換
func ConvertGRPCSubscriptionListResponse(iconList []*models.Subscription) []*subscription.Subscription {
	var subscriptions []*subscription.Subscription
	for _, v := range iconList {
		subscriptions = append(subscriptions, ConvertGRPCSubscriptionResponse(v))
	}
	return subscriptions
}
