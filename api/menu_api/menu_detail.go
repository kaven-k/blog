package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// MenuDetailView 菜单详情
func (MenuApi) MenuDetailView(c *gin.Context) {
	// 菜单的列表页是一个自定义的多对多关系，
	// 1.先查菜单（菜单不用分页）
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 需要菜单的ID，通过菜单ID去查第三张表，从第三张表中的banner_id取出关联的图片，然后进行排序
	// 2.查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	// 解决null值问题
	//banners := []Banner{}
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
	return
}
