package service

import (
	"context"

	"github.com/CA21engineer/Subs-server/apiServer/adopter"
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/golang/protobuf/ptypes"
)

// SubscriptionServiceImpl SubscriptionServiceImpl struct
type SubscriptionServiceImpl struct{}

// GetIconImageList サブスクを新規作成する際の追加可能アイコン一覧を取得する
func (SubscriptionServiceImpl) GetIconImageList(ctx context.Context, req *subscription.Empty) (*subscription.GetIconImageResponse, error) {
	icons, err := new(models.Icon).All()
	if err != nil {
		return nil, err
	}
	return &subscription.GetIconImageResponse{IconImage: adopter.ConvertGRPCIconListResponse(icons)}, nil
}

// GetSubscriptions サーバーに登録済みのサブスク一覧
func (SubscriptionServiceImpl) GetSubscriptions(context.Context, *subscription.Empty) (*subscription.GetSubscriptionsResponse, error) {
	subscriptions, err := new(models.Subscription).All()
	if err != nil {
		return nil, err
	}
	return &subscription.GetSubscriptionsResponse{Subscriptions: adopter.ConvertGRPCSubscriptionListResponse(subscriptions)}, nil
}

// GetPopularSubscriptions サーバーに登録済みのサブスク一覧を人気順で取得
func (SubscriptionServiceImpl) GetPopularSubscriptions(ctx context.Context, req *subscription.Empty) (*subscription.GetPopularSubscriptionsResponse, error) {
	return &subscription.GetPopularSubscriptionsResponse{}, nil
}

// GetRecommendSubscriptions サーバーに登録済みのサブスク一覧をパラメータによって出し分け
func (SubscriptionServiceImpl) GetRecommendSubscriptions(context.Context, *subscription.GetRecommendSubscriptionsRequest) (*subscription.GetRecommendSubscriptionsResponse, error) {
	return &subscription.GetRecommendSubscriptionsResponse{}, nil
}

// GetMySubscription 自分のリストに追加されているサブスク一覧
func (SubscriptionServiceImpl) GetMySubscription(ctx context.Context, req *subscription.GetMySubscriptionRequest) (*subscription.GetMySubscriptionResponse, error) {
	userSubscriptions, err := new(models.UserSubscription).GetUserSubscriptions(req.UserId)
	if err != nil {
		return nil, err
	}
	return &subscription.GetMySubscriptionResponse{Subscriptions: adopter.ConvertGRPCUserSubscriptionListResponse(userSubscriptions)}, nil
}

// CreateSubscription 未登録のサブスクを新規作成する
func (SubscriptionServiceImpl) CreateSubscription(ctx context.Context, req *subscription.CreateSubscriptionRequest) (*subscription.CreateSubscriptionResponse, error) {
	sub := models.NewSubscription(req.IconId, req.ServiceName, req.Price, req.Cycle, req.FreeTrial)
	startedAt, _ := ptypes.Timestamp(req.StartedAt)
	err := sub.Create(req.UserId, startedAt)
	if err != nil {
		return nil, err
	}
	return &subscription.CreateSubscriptionResponse{SubscriptionId: sub.SubscriptionID}, nil
}

// RegisterSubscription 登録済みのサブスクを自分のリストに追加する
func (SubscriptionServiceImpl) RegisterSubscription(context.Context, *subscription.RegisterSubscriptionRequest) (*subscription.RegisterSubscriptionResponse, error) {
	return &subscription.RegisterSubscriptionResponse{}, nil
}

// UpdateSubscription 既存サブスクを更新する
func (SubscriptionServiceImpl) UpdateSubscription(context.Context, *subscription.UpdateSubscriptionRequest) (*subscription.UpdateSubscriptionResponse, error) {
	return &subscription.UpdateSubscriptionResponse{}, nil
}

// UnregisterSubscription 登録済みのサブスクをリストから削除する
func (SubscriptionServiceImpl) UnregisterSubscription(context.Context, *subscription.UnregisterSubscriptionRequest) (*subscription.UnregisterSubscriptionResponse, error) {
	return &subscription.UnregisterSubscriptionResponse{}, nil
}
