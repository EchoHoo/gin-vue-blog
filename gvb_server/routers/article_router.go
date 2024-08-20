package routers

import (
	"gvb_server/gvb_server/api"
	"gvb_server/gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.DELETE("articles", app.ArticleRemoveView)
	router.GET("articles", app.ArticleListView) // 可以传入token 和 is_user 来控制用户发布的文章
	router.PUT("articles", app.ArticleUpdateView)

	router.GET("article/detail", app.ArticleDetailByTitleView)
	router.GET("article/:id", app.ArticleDetailView)
	router.GET("article/calendar", app.ArticleCalendarView)
	router.GET("article/tags", app.ArticleTagListView)

	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView) // 用户收藏的文章列表
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)

	router.GET("articles/text", app.FullTextSearchView)           //全文搜索
	router.GET("categorys", app.ArticleCategoryListView)          //全文搜索
	router.GET("article/content/:id", app.ArticleContentByIDView) //根据ID得到文章内容

	router.GET("article_id_title", app.ArticleIDTitleListView)

}
