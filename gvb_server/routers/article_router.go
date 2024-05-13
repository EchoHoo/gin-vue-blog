package routers

import (
	"gvb_server/gvb_server/api"
	"gvb_server/gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("article/detail", app.ArticleDetailByTitleView)
	router.GET("article/:id", app.ArticleDetailView)
	router.GET("article/calendar", app.ArticleCalendarView)
	router.GET("article/tags", app.ArticleTagListView)
	router.DELETE("articles", app.ArticleRemoveView)

	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.PUT("articles/collects", app.ArticleUpdateView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView)
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)

	router.GET("articles/text", app.FullTextSearchView) //全文搜索

}
