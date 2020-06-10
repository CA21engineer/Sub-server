package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/golang/protobuf/ptypes"
)

func ConvertGRPCUserSubscriptionResponse(s *models.UserSubscription) *subscription.Subscription {
	startedAt, _ := ptypes.TimestampProto(s.StartedAt)
	return &subscription.Subscription{
		SubscriptionId: s.Subscription.SubscriptionId,
		ServiceType:    s.Subscription.ServiceType,
		IconUri:        s.Subscription.Icon.IconUri,
		ServiceName:    s.Subscription.ServiceName,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      s.Subscription.FreeTrial,
		IsOriginal:     s.Subscription.IsOriginal,
		StartedAt:      startedAt,
	}
}

func ConvertGRPCUserSubscriptionListResponse(iconList []*models.UserSubscription) []*subscription.Subscription {
	var userSubscriptions []*subscription.Subscription
	for _, v := range iconList {
		userSubscriptions = append(userSubscriptions, ConvertGRPCUserSubscriptionResponse(v))
	}
	return userSubscriptions
}
