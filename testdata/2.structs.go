package main

import (
	"fmt"
	"github.com/fatih/structs"
	"server/models"
)

type AdvertisementRequest struct {
	models.MODEL `struct:"title"`
	Title        string `json:"title" binding:"required" msg:"请输入标题" struct:"title"`       // 显示的标题
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" struct:"-"`       // 跳转链接
	Images       string `json:"images" binding:"required,url" msg:"图片地址非法"`                // 图片
	IsShow       bool   `json:"is_show" binding:"required" msg:"请选择是否展示" struct:"is_show"` // 是否展示
}

func main() {
	u1 := AdvertisementRequest{
		Title:  "xxx",
		Href:   "xxx",
		Images: "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
