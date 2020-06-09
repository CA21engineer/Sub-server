package models

const (
	NOT_FOUND = iota + 1
  NOT_CATEGORIZED
  MUSIC
  MOVIE
  MATCHING
  STORAGE
)

type Subscription struct {
	SubscriptionId string
	IconId string
	ServiceName string
	ServiceType int32
	Price int32
	Cycle int32
	IsOriginal bool
	FreeTrial int32
}
