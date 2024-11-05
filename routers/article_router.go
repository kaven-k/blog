package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView) //
	router.GET("articles", app.ArticleListView)
}
