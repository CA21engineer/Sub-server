package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/golang/protobuf/ptypes"
)

// ConvertGRPCUserSubscriptionResponse `*models.UserSubscription`を`*subscription.Subscription`に変換
func ConvertGRPCUserSubscriptionResponse(s *models.UserSubscription) *subscription.Subscription {
	startedAt, _ := ptypes.TimestampProto(s.StartedAt)
	return &subscription.Subscription{
		SubscriptionId: s.SubscriptionID,
		ServiceType:    s.ServiceType,
		IconUri:        s.IconURI,
		ServiceName:    s.ServiceName,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      s.FreeTrial,
		IsOriginal:     s.IsOriginal,
		StartedAt:      startedAt,
	}
}

// ConvertGRPCUserSubscriptionListResponse `[]*models.UserSubscription`を`[]*subscription.Subscription`に変換
func ConvertGRPCUserSubscriptionListResponse(iconList []*models.UserSubscription) []*subscription.Subscription {
	var userSubscriptions []*subscription.Subscription
	for _, v := range iconList {
		userSubscriptions = append(userSubscriptions, ConvertGRPCUserSubscriptionResponse(v))
	}
	return userSubscriptions
}
