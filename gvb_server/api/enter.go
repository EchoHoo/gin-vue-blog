package api

import (
	advertapi "gvb_server/gvb_server/api/advert_api"
	"gvb_server/gvb_server/api/article_api"
	"gvb_server/gvb_server/api/chat_api"
	"gvb_server/gvb_server/api/comment_api"
	data_api "gvb_server/gvb_server/api/date_api"
	"gvb_server/gvb_server/api/digg_api"
	"gvb_server/gvb_server/api/images_api"
	"gvb_server/gvb_server/api/log_api"
	"gvb_server/gvb_server/api/menu_api"
	messageapi "gvb_server/gvb_server/api/message_api"
	"gvb_server/gvb_server/api/news_api"
	"gvb_server/gvb_server/api/settings_api"
	tagapi "gvb_server/gvb_server/api/tag_api"
	userapi "gvb_server/gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdverApi    advertapi.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     userapi.UserApi
	TagApi      tagapi.TagApi
	MessageApi  messageapi.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	NewsApi     news_api.NewApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
}

var ApiGroupApp = new(ApiGroup)
