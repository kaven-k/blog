package routers

import "server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("image_names", imagesApi.ImageNameListView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	router.PUT("images", imagesApi.ImageUpdateView)
}
