package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"` // 排序
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 切换的时间，单位是秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 切换的时间，单位是秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	// 参数校验
	var cr MenuRequest
	// 通过ShouldBindJSON方法尝试将客户端传来的 JSON 数据绑定到AdvertisementRequest结构体上
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 接下来就是添加，添加的时候主要有一个多表的操作（这个多表是自己生成的多表）
	// 重复值的判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}
	// 1.创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.FailWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		// 这里也需要判断 image_id 是否真的有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 2.给第三张表入库
	err = global.DB.Create(menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
