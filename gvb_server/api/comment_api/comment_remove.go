package comment_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/redis_ser"
	"gvb_server/gvb_server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}
	//统计评论下的字评论数
	subCommentList := FindSubCommentCount(commentModel)
	count := len(subCommentList) + 1
	redis_ser.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	//判断是否是字评论
	if commentModel.ParentCommentID != nil {
		//子评论
		//找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count -?", count))
	}
	//删除子评论以及根评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OKWithMessage(fmt.Sprintf("共 %d 条删除成功", len(deleteCommentIDList)), c)
}
