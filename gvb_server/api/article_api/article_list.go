package article_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"
	"gvb_server/gvb_server/service/redis_ser"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	boolSearch := elastic.NewBoolQuery()

	token := c.GetHeader("token")
	if cr.IsUser {
		claims, err := jwts.ParseToken(token)
		if err == nil && !redis_ser.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermQuery("user_id", claims.UserID))
		}
	}
	list, count, err := es_ser.CommList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content", "category"},
		Tag:      cr.Tag,
		Query:    boolSearch,
	})
	if err != nil {
		global.Log.Error(err)
		res.OKWithMessage("获取文章列表失败", c)
		return
	}
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MastMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}
	res.OkWithList(data, int64(count), c)
}
