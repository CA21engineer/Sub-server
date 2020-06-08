package service

import (
	"context"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

type SubscriptionServiceImpl struct{}

// サブスクを新規作成する際のアイコン一覧を取得する
func (SubscriptionServiceImpl) GetIconImageList(context.Context, *subscription.Empty) (*subscription.GetIconImageResponse, error) {
	return &subscription.GetIconImageResponse{}, nil
}

// サーバーに登録済みのサブスク一覧
func (SubscriptionServiceImpl) GetSubscription(context.Context, *subscription.GetSubscriptionRequest) (*subscription.GetSubscriptionResponse, error) {
	return &subscription.GetSubscriptionResponse{}, nil
}

// 自分のリストに追加されているサブスク一覧
func (SubscriptionServiceImpl) GetMySubscription(context.Context, *subscription.GetMySubscriptionRequest) (*subscription.GetMySubscriptionResponse, error) {
	return &subscription.GetMySubscriptionResponse{}, nil
}

// 未登録のサブスクを新規作成する
func (SubscriptionServiceImpl) CreateSubscription(context.Context, *subscription.CreateSubscriptionRequest) (*subscription.CreateSubscriptionResponse, error) {
	return &subscription.CreateSubscriptionResponse{}, nil
}

// 登録済みのサブスクを自分のリストに追加する
func (SubscriptionServiceImpl) RegisterSubscription(context.Context, *subscription.RegisterSubscriptionRequest) (*subscription.RegisterSubscriptionResponse, error) {
	return &subscription.RegisterSubscriptionResponse{}, nil
}
