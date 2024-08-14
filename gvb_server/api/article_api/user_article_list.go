package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) UserArticleListView(c *gin.Context) {
	var cr models.PageInfo
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	from := (cr.Page - 1) * cr.Limit
	boolSearch := elastic.NewTermsQuery("user_id", claims.UserID)
	if cr.Page == 0 {
		cr.Page = 1
	}
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	var list = make([]models.ArticleModel, 0)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From(from).
		Size(cr.Limit).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	count := result.Hits.TotalHits.Value
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err := json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue

		}

		article.ID = hit.Id
		list = append(list, article)
	}
	res.OkWithList(list, count, c)
}
