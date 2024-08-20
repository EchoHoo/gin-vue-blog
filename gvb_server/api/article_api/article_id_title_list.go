package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"

	"github.com/olivere/elastic/v7"
)

type ArticleIDTitleListResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (ArticleApi) ArticleIDTitleListView(c *gin.Context) {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Source(`{"_source":["title"]}`).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取文章列表失败", c)
		return
	}
	var articleIDTitleList []ArticleIDTitleListResponse
	for _, hit := range result.Hits.Hits {
		model := models.ArticleModel{}
		json.Unmarshal(hit.Source, &model)
		model.ID = hit.Id
		articleIDTitleList = append(articleIDTitleList, ArticleIDTitleListResponse{
			Value: model.ID,
			Label: model.Title,
		})
	}
	res.OkWithData(articleIDTitleList, c)
}
