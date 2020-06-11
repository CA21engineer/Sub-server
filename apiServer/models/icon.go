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

// All 追加可能アイコン一覧を返す
func (i *Icon) All() ([]*Icon, error) {
	var icons []*Icon
	if err := DB.Where("is_original = ?", true).Find(&icons).Error; err != nil {
		return nil, err
	}
	return icons, nil
}
