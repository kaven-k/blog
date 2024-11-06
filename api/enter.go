//

package api

import (
	"server/api/advertisement_api"
	"server/api/article_api"
	"server/api/digg_api"
	"server/api/images_api"
	"server/api/menu_api"
	"server/api/message_api"
	"server/api/settings_api"
	"server/api/tag_api"
	"server/api/user_api"
)

type ApiGroup struct {
	SettingsApi      settings_api.SettingsApi
	ImagesApi        images_api.ImagesApi
	AdvertisementApi advertisement_api.AdvertisementApi
	MenuApi          menu_api.MenuApi
	UserApi          user_api.UserApi
	TagApi           tag_api.TagApi
	MessageApi       message_api.MessageApi
	ArticleApi       article_api.ArticleApi
	DiggApi          digg_api.DiggApi
}

// ApiGroupApp 实例化这个对象
var ApiGroupApp = new(ApiGroup)
