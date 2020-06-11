package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// ConvertGRPCSubscriptionResponse `*models.Subscription`を`*subscription.Subscription`に変換
func ConvertGRPCSubscriptionResponse(s *models.SubscriptionWithIcon) *subscription.Subscription {
	return &subscription.Subscription{
		SubscriptionId: s.Subscription.SubscriptionID,
		ServiceType:    s.Subscription.ServiceType,
		IconUri:        s.Icon.IconURI,
		ServiceName:    s.Subscription.ServiceName,
		Price:          s.Subscription.Price,
		Cycle:          s.Subscription.Cycle,
		FreeTrial:      s.Subscription.FreeTrial,
		IsOriginal:     s.Subscription.IsOriginal,
		StartedAt:      nil,
	}
}

// ConvertGRPCSubscriptionListResponse `[]*models.Subscription`を`[]*subscription.Subscription`に変換
func ConvertGRPCSubscriptionListResponse(iconList []*models.SubscriptionWithIcon) []*subscription.Subscription {
	var subscriptions []*subscription.Subscription
	for _, v := range iconList {
		subscriptions = append(subscriptions, ConvertGRPCSubscriptionResponse(v))
	}
	return subscriptions
}
