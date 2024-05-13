package article_api

import (
	"context"
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 如果文章删除了，用户收藏这篇文章怎么办？
	// 1.顺带把文章关联的收藏也删除了
	// 2.用户收藏表，新增一个字段，记录文章是否删除，如果删除了，就不再显示在收藏列表中
	// 3.类似redis的惰性删除

	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		go es_ser.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OKWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)

}
