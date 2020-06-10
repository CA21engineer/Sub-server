package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/golang/protobuf/ptypes"
)

func ConvertGRPCSubscriptionResponse(s *models.Subscription) *subscription.Subscription {
	startedAt, _ := ptypes.TimestampProto(s.StartedAt)
	return &subscription.Subscription{
		SubscriptionId: s.SubscriptionId,
		ServiceType:    s.ServiceType,
		IconUri:        s.Icon.IconUri,
		ServiceName:    s.ServiceName,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      int32(s.FreeTrial.Day()),
		IsOriginal:     s.IsOriginal,
		StartedAt:      startedAt,
	}
}

func ConvertGRPCSubscriptionListResponse(iconList []*models.Subscription) []*subscription.Subscription {
	var subscriptions []*subscription.Subscription
	for _, v := range iconList {
		subscriptions = append(subscriptions, ConvertGRPCSubscriptionResponse(v))
	}
	return subscriptions
}
