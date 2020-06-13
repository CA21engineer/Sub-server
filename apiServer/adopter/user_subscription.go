package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// ConvertGRPCUserSubscriptionResponse `*models.UserSubscription`を`*subscription.UserSubscription`に変換
func ConvertGRPCUserSubscriptionResponse(s *models.UserSubscription) *subscription.UserSubscription {
	subscriptionWithIcon := models.SubscriptionWithIcon{
		Subscription: s.Subscription,
	}
	return &subscription.UserSubscription{
		UserSubscriptionId:s.UserSubscriptionID,
		Subscription: ConvertGRPCSubscriptionResponse(&subscriptionWithIcon),
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
