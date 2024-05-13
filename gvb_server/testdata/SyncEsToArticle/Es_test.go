package testdata

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/gvb_server/core"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/service/es_ser"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSearch(t *testing.T) {
	core.InitConf()

	core.InitLogger()
	global.ESClient = core.EsConnect()

	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		indexList := es_ser.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		bulk := global.ESClient.Bulk()
		for _, id := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(id)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		fmt.Println(article.Title, "添加成功 ", len(indexList), "条数据", result, err)
	}
}

func TestAsyncSearch(t *testing.T) {
	core.InitConf()
	core.InitLogger()
	global.ESClient = core.EsConnect()

}
