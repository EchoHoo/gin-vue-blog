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

func (ArticleApi) ArticleTagListView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	// 获取总计数
	countAgg := elastic.NewCardinalityAggregation().Field("tags.keyword")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags_count", countAgg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error("Elasticsearch query failed", err)
		res.FailWithMessage("Elasticsearch query failed: "+err.Error(), c)
		return
	}

	cTag, found := result.Aggregations.Cardinality("tags_count")
	if !found || cTag == nil {
		global.Log.Error("Cardinality aggregation 'tags_count' not found")
		res.FailWithMessage("Internal error", c)
		return
	}
	count := int64(*cTag.Value)

	// 执行第二次查询
	agg := elastic.NewTermsAggregation().Field("tags.keyword")
	agg.SubAggregation("articles", elastic.NewTopHitsAggregation().Size(100).FetchSourceContext(elastic.NewFetchSourceContext(true).Include("abstract")))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error("Elasticsearch query failed", err)
		res.FailWithMessage("Elasticsearch query failed: "+err.Error(), c)
		return
	}

	rawTags, found := result.Aggregations["tags"]
	if !found {
		global.Log.Error("tags aggregation not found")
		res.FailWithMessage("Internal error", c)
		return
	}

	// 输出原始结果以进行调试
	global.Log.Info("Raw tags aggregation result: ", string(rawTags))

	var tagType TagsType
	err = json.Unmarshal(rawTags, &tagType)
	if err != nil {
		global.Log.Error("Failed to unmarshal tags aggregation", err)
		res.FailWithMessage("Internal error", c)
		return
	}

	var tagList []TagsResponse
	for _, bucket := range tagType.Buckets {
		var articleList []string
		for _, hit := range bucket.Articles.Hits.Hits {
			var article struct {
				Abstract string `json:"abstract"`
			}
			if err := json.Unmarshal(hit.Source, &article); err != nil {
				global.Log.Error("Failed to unmarshal top hits", err)
				continue
			}
			articleList = append(articleList, article.Abstract)
		}
		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}

	res.OkWithList(tagList, count, c)
}

type TagsType struct {
	Buckets []struct {
		Key      string `json:"key"`
		DocCount int64  `json:"doc_count"`
		Articles struct {
			Hits struct {
				Hits []struct {
					Source json.RawMessage `json:"_source"`
				} `json:"hits"`
			} `json:"hits"`
		} `json:"articles"`
	} `json:"buckets"`
}

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int64    `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
}
