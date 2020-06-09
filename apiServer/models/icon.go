package models

type Icon struct {
	IconId string `gorm:`
	IconUri string
	IsOriginal bool
}

type Icons []*Icon


func NewIcon(id, uri string, is_original bool) *Icon {
	return &Icon{id, uri, is_original}
}

// All 全てのiconを取得
func (i *Icon) All()([]*Icon, error){
	var icons Icons
	if err := DB.Find(&icons).Error; err != nil {
		return nil, err
	}
	return icons, nil
}
