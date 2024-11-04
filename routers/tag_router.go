package routers

import "server/api"

func (router RouterGroup) TagRouter() {
	tag := api.ApiGroupApp.TagApi
	router.POST("tags", tag.TagCreateView)
	router.GET("tags", tag.TagListView)
	router.PUT("tags/:id", tag.TagUpdateView)
	router.DELETE("tags", tag.TagRemoveView)

}
