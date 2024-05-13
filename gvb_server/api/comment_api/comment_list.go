package comment_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	fmt.Println("cr.ArticleID = ", cr.ArticleID)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)

	res.OkWithData(filter.Select("c", rootCommentList), c)
	// res.OkWithData(filter.Omit("c", rootCommentList), c)
}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	global.DB.Preload("User").Find(&RootCommentList, "article_id =? and parent_comment_id is null", articleID)
	diggInfo := redis_ser.NewCommentDigg().GetInfo()
	for _, model := range RootCommentList {
		var subCommentList, newSubCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)

		for _, commentModel := range subCommentList {
			digg := diggInfo[fmt.Sprintf("%d", commentModel.ID)]
			commentModel.DiggCount = commentModel.DiggCount + int(digg)
			newSubCommentList = append(newSubCommentList, commentModel)
		}
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount = model.DiggCount + modelDigg
		model.SubComments = newSubCommentList
	}

	return RootCommentList
}

// 递归查看子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {

		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}

// 得到子评论列表
func FindSubCommentCount(model models.CommentModel) (subCommentList []models.CommentModel) {

	findSubCommentList(model, &subCommentList)
	return subCommentList
}

// 递归查看子列表
func findSubCommentList(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {

		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}
