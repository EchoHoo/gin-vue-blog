package esarticletag

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/gvb_server/core"
	"gvb_server/gvb_server/flag"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"testing"

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

func ArticleTagListView() {
	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	query := elastic.NewBoolQuery()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	//fmt.Println(string(result.Aggregations["tags"]))
	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
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
	fmt.Println(tagList)
}

func TestTags(t *testing.T) {
	//读取配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//连接redis
	global.Redis = core.ConnectRedis()

	//连接ES
	global.ESClient = core.EsConnect()

	ArticleTagListView()

}

// {"doc_count_error_upper_bound":0,
// "sum_other_doc_count":0,
// "buckets":[
// 	{"key":"python","doc_count":1,"articles":
// 		{"doc_count_error_upper_bound":0,
// 		"sum_other_doc_count":0,
// 		"buckets":[{"key":"b","doc_count":1}]}
// 	},
// 	{"key":"后端","doc_count":1,"articles":
// 		{"doc_count_error_upper_bound":0,
// 		"sum_other_doc_count":0,
// 		"buckets":[{"key":"b","doc_count":1}]}}
// 		]
// 	}
