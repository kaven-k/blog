package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), app.MessageCreateView)
	router.GET("messages_list_admin", middleware.JwtAdmin(), app.MessageListAdminView)
	router.GET("messages_list_user", middleware.JwtAuth(), app.MessageListUserView)
	router.GET("messages_record", middleware.JwtAuth(), app.MessageRecord)
}
