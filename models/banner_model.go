package models

import (
	"gorm.io/gorm"
	"os"
	"server/global"
	"server/models/ctype"
)

type BannerModel struct {
	MODEL
	Path          string              `json:"path"`                            // 图片路径
	Hash          string              `json:"hash"`                            // 图片的hash值，用于判断重复图片
	Name          string              `gorm:"size:38" json:"name"`             // 图片名称
	ImageLocation ctype.ImageLocation `gorm:"default:1" json:"image_location"` // 标记一个标志，照片存储的位置
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageLocation == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
