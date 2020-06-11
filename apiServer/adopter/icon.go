package adopter

import (
	"github.com/CA21engineer/Subs-server/apiServer/models"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
)

// ConvertGRPCIconResponse `*model.Icon`を`*subscription.IconImage`に変換
func ConvertGRPCIconResponse(i *models.Icon) *subscription.IconImage {
	return &subscription.IconImage{
		IconId:  i.IconID,
		IconUri: i.IconURI,
	}
}

// ConvertGRPCIconListResponse `[]*model.Icon`を`[]*subscription.IconImage`に変換
func ConvertGRPCIconListResponse(iconList []*models.Icon) []*subscription.IconImage {
	var iconImages []*subscription.IconImage
	for _, v := range iconList {
		iconImages = append(iconImages, ConvertGRPCIconResponse(v))
	}
	return iconImages
}
