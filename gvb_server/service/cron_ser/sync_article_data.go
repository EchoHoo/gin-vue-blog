package cron_ser

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/service/redis_ser"

	"github.com/olivere/elastic/v7"
)

func SyncArticleDate() {
	// 1.查询es中的全部数据，为后面的数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error("sync article data failed", err)
		return
	}
	//2. 拿到redis中的缓存数据
	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error("sync article data failed", err)
			continue
		}
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		//3.计算新的数据
		newDigg := article.DiggCount + int(digg)
		newLook := article.LookCount + int(look)
		newComment := article.CommentCount + int(comment)

		if digg == 0 && look == 0 && comment == 0 {
			//global.Log.Info("sync article data no change", article.Title)
			continue
		}
		//更新
		global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).Do(context.Background())

		global.Log.Info("sync article data success ArticleTitle: ", article.Title)
	}
	//清除redis的缓存数据
	redis_ser.NewDigg().Clear()
	redis_ser.NewLook().Clear()
	redis_ser.NewCommentCount().Clear()
}
