package routers

import "gvb_server/gvb_server/api"

func (router RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	router.POST("news", app.NewListView)

}
