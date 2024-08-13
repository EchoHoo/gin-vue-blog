package routers

import (
	"gvb_server/gvb_server/api"
	"gvb_server/gvb_server/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("GuoChengHao666DaShuaiGe"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi

	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAuth(), app.UserUpdateRoleView)
	router.POST("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.PUT("logout", middleware.JwtAuth(), app.UserLogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("users", app.UserCreateView)
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)
	router.PUT("user_info", middleware.JwtAuth(), app.UserUpdateNicknameView)
}
