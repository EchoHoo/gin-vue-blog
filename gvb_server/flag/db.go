package flag

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/plugins/log_stash"
)

// 表的迁移
func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{}) //ok
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})         //ok

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},  //ok
			&models.TagModel{},     //ok
			&models.MessageModel{}, //ok
			&models.AdvertModel{},  //ok
			&models.UserModel{},    //ok
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&log_stash.LogStashModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据表结构失败 ")
		return
	}
	global.Log.Info("[success] 生成数据库表结构成功!")
}
