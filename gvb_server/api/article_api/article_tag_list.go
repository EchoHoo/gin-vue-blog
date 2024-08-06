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

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

// func (ArticleApi) ArticleTagListView(c *gin.Context) {

// 	var cr models.PageInfo
// 	_ = c.ShouldBindQuery(&cr)

// 	if cr.Limit == 0 {
// 		cr.Limit = 10
// 	}
// 	offset := (cr.Page - 1) * cr.Limit
// 	if offset < 0 {
// 		offset = 0
// 	}

// 	result, err := global.ESClient.
// 		Search(models.ArticleModel{}.Index()).
// 		Aggregation("tags", elastic.NewValueCountAggregation().Field("tags.keyword")).
// 		Size(0).
// 		Do(context.Background())
// 	cTag, _ := result.Aggregations.Cardinality("tags")
// 	count := int64(*cTag.Value)

// 	agg := elastic.NewTermsAggregation().Field("tags")

// 	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
// 	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

// 	query := elastic.NewBoolQuery()

// 	result, err = global.ESClient.
// 		Search(models.ArticleModel{}.Index()).
// 		Query(query).
// 		Aggregation("tags", agg).
// 		Size(0).
// 		Do(context.Background())
// 	if err != nil {
// 		global.Log.Error(err)
// 		res.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	var tagType TagsType
// 	var tagList = make([]TagsResponse, 0)
// 	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
// 	for _, bucket := range tagType.Buckets {

// 		var articleList []string
// 		for _, s := range bucket.Articles.Buckets {
// 			articleList = append(articleList, s.Key)
// 		}

// 		tagList = append(tagList, TagsResponse{
// 			Tag:           bucket.Key,
// 			Count:         bucket.DocCount,
// 			ArticleIDList: articleList,
// 		})
// 	}

//		res.OkWithList(tagList, count, c)
//	}
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

	// 打印原始数据
	printOriginalData()

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
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("_id.keyword"))
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
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}
		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}

	res.OkWithList(tagList, count, c)
}

// func (ArticleApi) ArticleTagListView(c *gin.Context) {
// 	var cr models.PageInfo
// 	_ = c.ShouldBindQuery(&cr)

// 	if cr.Limit == 0 {
// 		cr.Limit = 10
// 	}
// 	offset := (cr.Page - 1) * cr.Limit
// 	if offset < 0 {
// 		offset = 0
// 	}

// 	// 第一次查询，获取总计数
// 	countAgg := elastic.NewCardinalityAggregation().Field("tags.keyword")
// 	result, err := global.ESClient.
// 		Search(models.ArticleModel{}.Index()).
// 		Aggregation("tags_count", countAgg).
// 		Size(0).
// 		Do(context.Background())

// 	if err != nil {
// 		global.Log.Error("Elasticsearch query failed", err)
// 		res.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	tagsCount, found := result.Aggregations.Cardinality("tags_count")
// 	if !found {
// 		global.Log.Error("tags_count aggregation not found")
// 		res.FailWithMessage("Internal error", c)
// 		return
// 	}
// 	count := int64(*tagsCount.Value)

// 	// 第二次查询，获取分页结果
// 	termsAgg := elastic.NewTermsAggregation().Field("tags.keyword").Size(1000)
// 	// termsAgg.SubAggregation("articles", elastic.NewTermsAggregation().Field("_id.keyword"))
// 	termsAgg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
// 	termsAgg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

// 	query := elastic.NewBoolQuery()

// 	result, err = global.ESClient.
// 		Search(models.ArticleModel{}.Index()).
// 		Query(query).
// 		Aggregation("tags", termsAgg).
// 		Size(0).
// 		Do(context.Background())

// 	if err != nil {
// 		global.Log.Error("Elasticsearch query failed", err)
// 		res.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	// 解析结果
// 	rawTags, found := result.Aggregations["tags"]
// 	if !found {
// 		global.Log.Error("tags aggregation not found")
// 		res.FailWithMessage("Internal error", c)
// 		return
// 	}

// 	global.Log.Info("Raw tags aggregation result: ", string(rawTags))

// 	var tagsAgg TagsType
// 	err = json.Unmarshal(rawTags, &tagsAgg)
// 	jsonData, _ := json.MarshalIndent(tagsAgg, "", "  ")
// 	global.Log.Info("JSON tags aggregation result: ", string(jsonData))

// 	if err != nil {
// 		global.Log.Error("Failed to unmarshal tags aggregation", err)
// 		res.FailWithMessage("Internal error", c)
// 		return
// 	}

// 	global.Log.Info("Parsed tags aggregation: ", tagsAgg)

// 	var tagList []TagsResponse
// 	for _, bucket := range tagsAgg.Buckets {
// 		var articleList []string
// 		for _, articleBucket := range bucket.Articles.Buckets {
// 			global.Log.Info("Article ID: ", articleBucket.Key) // 输出文章ID以调试
// 			articleList = append(articleList, articleBucket.Key)
// 		}

// 		tagList = append(tagList, TagsResponse{
// 			Tag:           bucket.Key,
// 			Count:         bucket.DocCount,
// 			ArticleIDList: articleList,
// 		})
// 	}

// 	res.OkWithList(tagList, count, c)
// }

func printOriginalData() {
	// 执行查询并打印数据
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Size(10). // 可以根据需要调整大小
		Do(context.Background())
	if err != nil {
		global.Log.Error("Elasticsearch query failed", err)
		return
	}

	// 打印查询结果
	for _, hit := range result.Hits.Hits {
		global.Log.Info("Document ID: ", hit.Id)
		global.Log.Info("Document Source: ", string(hit.Source))
	}
}
