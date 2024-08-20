package comment_api

import (
	"context"
	"encoding/json"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CommonListResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	ArticleTitle    string    `json:"article_title"`
	ArticleBanner   string    `json:"article_banner"`
	ParentCommentID *uint     `json:"parent_comment_id"`
	Content         string    `json:"content"`
	DiggCount       int       `json:"digg_count"`
	CommentCount    int       `json:"comment_count"`
	UserNickName    string    `json:"user_nick_name"`
}

func (CommentApi) CommentListAllView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.CommentModel{}, common.Option{
		PageInfo: cr,
		Preload:  []string{"User"},
	})
	var commentList = make([]CommonListResponse, 0)
	var articleIDList []interface{}
	var collMap = map[string]models.ArticleModel{}
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = models.ArticleModel{}
	}
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error("json unmarshal error", err)
			continue
		}
		collMap[hit.Id] = article
	}
	for _, model := range list {
		commentList = append(commentList, CommonListResponse{
			ID:              model.ID,
			CreatedAt:       model.CreateAt,
			ParentCommentID: model.ParentCommentID,
			Content:         model.Content,
			DiggCount:       model.DiggCount,
			CommentCount:    model.CommentCount,
			UserNickName:    model.User.NickName,
			ArticleTitle:    collMap[model.ArticleID].Title,
			ArticleBanner:   collMap[model.ArticleID].BannerUrl,
		})
	}
	res.OkWithList(commentList, count, c)
}
