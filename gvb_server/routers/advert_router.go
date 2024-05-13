package routers

import "gvb_server/gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdverApi
	router.POST("adverts", app.AdvertCreateView)
	router.GET("adverts", app.AdvertListView)
	router.PUT("adverts/:id", app.AdvertUpdateView)
	router.DELETE("adverts", app.AdvertRemoveView)
}
