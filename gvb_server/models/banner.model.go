package models

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片的类型，本地还是七牛
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
		}
	}
	return nil
}
