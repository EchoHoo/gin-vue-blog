package routers

import (
	"gvb_server/gvb_server/api"
)

//	func SettingsRouter(router *gin.Engine) {
//		settingsApi := api.ApiGroupApp.SettingsApi
//		router.GET("", settingsApi.SettingsInfoView)
//	}
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// router.GET("settings", settingsApi.SettingsInfoView)
	// router.PUT("settings", settingsApi.SettingsInfoUpdateView)
	// router.GET("settings_email", settingsApi.SettingsEmailView)
	// router.PUT("settings_email", settingsApi.SettingsEmailInfoUpdateView)
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
