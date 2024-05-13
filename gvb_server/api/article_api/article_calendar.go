package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	//时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day") //这里填的就是聚合的维度

	//时间段搜索
	//从今天开始，到去年的今天
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0) //一年前
	format := "2006-01-02 15:04:05"
	//Less Than or Great Than
	query := elastic.NewRangeQuery("created_at").Gte(aYearAgo.Format(format)).Lte(now.Format(format))

	result, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}
	//var resList = make([]CalendarResponse, 0)
	// days := int(now.Sub(aYearAgo).Hours() / 24)
	// for i := 0; i < days; i++ {

	// }
	var data BucketsType
	json.Unmarshal(result.Aggregations["calendar"], &data)

	var resList = make([]CalendarResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		resList = append(resList, CalendarResponse{
			Date:  Time.Format("2006-01-02"),
			Count: bucket.DocCount,
		})
	}
	res.OkWithData(resList, c)

}
