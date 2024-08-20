package routers

import (
	"gvb_server/gvb_server/api"
	"gvb_server/gvb_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments", app.CommentListView)
	router.GET("comments_all", app.CommentListAllView)
	router.GET("comments/:id", app.CommentDigg)
	router.DELETE("comments/:id", app.CommentRemoveView)
}
