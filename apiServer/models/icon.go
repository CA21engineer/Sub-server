package models

// Icon struct
type Icon struct {
	IconID     string
	IconURI    string
	IsOriginal bool
}

// NewIcon New Icon struct
func NewIcon(id, uri string, isOriginal bool) *Icon {
	return &Icon{id, uri, isOriginal}
}

// All 全てのiconを取得
func (i *Icon) All() ([]*Icon, error) {
	var icons []*Icon
	if err := DB.Find(&icons).Error; err != nil {
		return nil, err
	}
	return icons, nil
}
