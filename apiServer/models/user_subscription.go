package models

type UserSubscription struct {
	UserSubscriptionId string
	UserId             string
	SubscriptionId     string
	Cycle              int32
	Price              int32
}
