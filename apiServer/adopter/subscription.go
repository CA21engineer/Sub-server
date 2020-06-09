package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

func ConvertGRPCSubscriptionResponse(s *models.Subscription) *subscription.Subscription {
	return &subscription.Subscription{
		SubscriptionId: s.SubscriptionId,
		ServiceType:    s.IconId,
		IconUri:        s.ServiceName,
		ServiceName:    s.ServiceType,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      s.FreeTrial,
		IsOriginal:     s.IsOriginal,
		StartedAt:      s.StartedAt,
	}
}

func ConvertGRPCSubscriptionListResponse(iconList []*models.Subscription) []*subscription.Subscription {
	var subscriptions []*subscription.Subscription
	for _, v := range iconList {
		subscriptions = append(subscriptions, ConvertGRPCSubscriptionResponse(v))
	}
	return subscriptions
}
