package article_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"
	"gvb_server/gvb_server/service/redis_ser"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	IsCollect bool `json:"is_collect" form:"is_collect"` // 是否收藏
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_ser.NewLook().Set(cr.ID)
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	articleDetail := ArticleDetailResponse{
		ArticleModel: model,
		IsCollect:    IsUserArticleCollect(c, cr.ID),
	}

	//　去查一下用户是否收藏文章
	res.OkWithData(articleDetail, c)
}
func IsUserArticleCollect(c *gin.Context, articleID string) bool {
	token := c.Request.Header.Get("token")
	if token == "" {
		return false
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return false
	}
	//判断是否在redis中存在，存在就是注销过了
	if redis_ser.CheckLogout(token) {
		return false
	}
	var count int64
	global.DB.Model(models.UserCollectModel{}).Where("user_id = ? and article_id = ?", claims.UserID, articleID).Count(&count)

	return count > 0
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
