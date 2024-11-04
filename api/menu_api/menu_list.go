package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 列表页
func (MenuApi) MenuListView(c *gin.Context) {
	// 菜单的列表页是一个自定义的多对多关系，
	// 1.先查菜单（菜单不用分页）
	var menuList []models.MenuModel
	var menuIDLIst []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDLIst)
	// 需要菜单的ID，通过菜单ID去查第三张表，从第三张表中的banner_id取出关联的图片，然后进行排序
	// 2.查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDLIst)
	var menus []MenuResponse
	for _, model := range menuList {
		// model就是一个菜单
		// 解决null值问题
		//banners := []Banner{}
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model, // 使用继承的好处就是不需要去写额外的字段
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
