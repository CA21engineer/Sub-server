package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/golang/protobuf/ptypes"
)

// ConvertGRPCUserSubscriptionResponse `*models.UserSubscription`を`*subscription.UserSubscription`に変換
func ConvertGRPCUserSubscriptionResponse(s *models.UserSubscription) *subscription.UserSubscription {
	startedAt, _ := ptypes.TimestampProto(s.StartedAt)
	return &subscription.UserSubscription{
		UserSubscriptionId:s.UserSubscriptionID,
		SubscriptionId: s.Subscription.SubscriptionID,
		ServiceType:    s.Subscription.ServiceType,
		IconUri:        s.Icon.IconURI,
		ServiceName:    s.Subscription.ServiceName,
		Price:          s.Price,
		Cycle:          s.Cycle,
		FreeTrial:      s.Subscription.FreeTrial,
		IsOriginal:     s.Subscription.IsOriginal,
		StartedAt:      startedAt,
	}
}

// ConvertGRPCUserSubscriptionListResponse `[]*models.UserSubscription`を`[]*subscription.UserSubscription`に変換
func ConvertGRPCUserSubscriptionListResponse(iconList []*models.UserSubscription) []*subscription.UserSubscription {
	var userSubscriptions []*subscription.UserSubscription
	for _, v := range iconList {
		userSubscriptions = append(userSubscriptions, ConvertGRPCUserSubscriptionResponse(v))
	}
	return userSubscriptions
}
