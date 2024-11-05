package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView) //
	router.GET("articles", app.ArticleListView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles/update", app.ArticleUpdateView)
	router.GET("articles/:id", app.ArticleDetailView)

}
