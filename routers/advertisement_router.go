package routers

import "server/api"

func (router RouterGroup) AdvertisementRouter() {
	advertisement := api.ApiGroupApp.AdvertisementApi
	router.POST("ad", advertisement.AdvertisementCreateView)
	router.GET("ad", advertisement.AdvertisementListView)
	router.PUT("ad/:id", advertisement.AdvertisementUpdateView)
	router.DELETE("ad", advertisement.AdvertisementRemoveView)

}
