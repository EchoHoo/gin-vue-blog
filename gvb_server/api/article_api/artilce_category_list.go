package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type CategoryResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		}
	}
	agg := elastic.NewTermsAggregation().Field("category")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("category", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["category"]
	var CategoryType T
	json.Unmarshal(byteData, &CategoryType)
	var categoryList = make([]CategoryResponse, 0)
	for _, v := range CategoryType.Buckets {	
		if v.Key == ""{
			continue
		}
		categoryList = append(categoryList, CategoryResponse{
			Label: v.Key,
			Value: v.Key,
		})
	}
	res.OkWithData(categoryList, c)
}
