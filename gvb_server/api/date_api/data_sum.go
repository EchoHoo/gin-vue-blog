package data_api

import (
	"context"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	CharGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
}

func (DataApi) DataSumView(c *gin.Context) {
	var userCount, articleCount, messageCount, chatGroupCount int
	var nowLoginCount, nowSignCount int
	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value)
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&chatGroupCount)
	global.DB.Model(models.LoginDataModel{}).Where("to_days(create_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(create_at)=to_days(now())").
		Select("count(id)").Scan(&nowSignCount)

	res.OkWithData(DataSumResponse{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		CharGroupCount: chatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}
