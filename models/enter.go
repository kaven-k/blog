package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id" structs:"-"` // 主键ID
	CreatedAt time.Time `json:"created_at" structs:"-"`           // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                    // 更新时间
}

type PageInfo struct {
	Page  int    `form:"page" json:"page"`   // 当前请求的页码
	Key   string `form:"key" json:"key"`     // 关键字
	Limit int    `form:"limit" json:"limit"` // 每页显示的数据数量
	Sort  string `form:"sort" json:"sort"`   // 用于数据的排序
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type EsIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}
