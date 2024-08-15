package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	global.DB.Debug().Find(&collects, "user_id =? AND article_id IN (?)", claims.UserID, cr.IDList).
		Select("article_id").
		Scan(&articleIDList)

	fmt.Println(articleIDList)

	if len(articleIDList) == 0 {
		res.FailWithMessage("未收藏任何文章", c)
		return
	}
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		count := article.CollectsCount - 1
		err = es_ser.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}

	}
	global.DB.Where("article_id IN (?)", idList).Delete(&collects)
	res.OKWithMessage(fmt.Sprintf("删除成功，共删除%d篇文章", len(articleIDList)), c)
}
