package service

import (
	"context"

	"github.com/CA21engineer/Subs-server/apiServer/adopter"
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

type SubscriptionServiceImpl struct{}

// サブスクを新規作成する際のアイコン一覧を取得する
func (SubscriptionServiceImpl) GetIconImageList(ctx context.Context, req *subscription.Empty) (*subscription.GetIconImageResponse, error) {
	icons, err := new(models.Icon).All()
	if err != nil {
		return nil, err
	}
	return &subscription.GetIconImageResponse{IconImage: adopter.ConvertGRPCIconListResponse(icons)}, nil
}

// サーバーに登録済みのサブスク一覧
func (SubscriptionServiceImpl) GetSubscriptions(context.Context, *subscription.GetSubscriptionsRequest) (*subscription.GetSubscriptionsResponse, error) {
	return &subscription.GetSubscriptionsResponse{}, nil
}

// 自分のリストに追加されているサブスク一覧
func (SubscriptionServiceImpl) GetMySubscription(ctx context.Context, req *subscription.GetMySubscriptionRequest) (*subscription.GetMySubscriptionResponse, error) {
	userSubscriptions, err := new(models.UserSubscription).GetUserSubscriptions(req.UserId)
	if err != nil {
		return nil, err
	}
	return &subscription.GetMySubscriptionResponse{Subscriptions: adopter.ConvertGRPCUserSubscriptionListResponse(userSubscriptions)}, nil
}

// 未登録のサブスクを新規作成する
func (SubscriptionServiceImpl) CreateSubscription(context.Context, *subscription.CreateSubscriptionRequest) (*subscription.CreateSubscriptionResponse, error) {
	return &subscription.CreateSubscriptionResponse{}, nil
}

// 登録済みのサブスクを自分のリストに追加する
func (SubscriptionServiceImpl) RegisterSubscription(context.Context, *subscription.RegisterSubscriptionRequest) (*subscription.RegisterSubscriptionResponse, error) {
	return &subscription.RegisterSubscriptionResponse{}, nil
}

// 既存サブスクを更新する
func (SubscriptionServiceImpl) UpdateSubscription(context.Context, *subscription.UpdateSubscriptionRequest) (*subscription.UpdateSubscriptionResponse, error) {
	return &subscription.UpdateSubscriptionResponse{}, nil
}

// 登録済みのサブスクをリストから削除する
func (SubscriptionServiceImpl) UnregisterSubscription(context.Context, *subscription.UnregisterSubscriptionRequest) (*subscription.UnregisterSubscriptionResponse, error) {
	return &subscription.UnregisterSubscriptionResponse{}, nil
}
