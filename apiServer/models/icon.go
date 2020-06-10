package models

import "time"

// Icon struct
type Icon struct {
	IconID     string
	IconURI    string
	IsOriginal bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewIcon New Icon struct
func NewIcon(id, uri string, isOriginal bool) *Icon {
	return &Icon{id, uri, isOriginal, time.Now(), time.Now()}
}

// All 全てのiconを取得
func (i *Icon) All() ([]*Icon, error) {
	var icons []*Icon
	if err := DB.Find(&icons).Error; err != nil {
		return nil, err
	}
	return icons, nil
}
