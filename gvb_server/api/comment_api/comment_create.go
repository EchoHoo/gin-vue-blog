package comment_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"
	"gvb_server/gvb_server/service/redis_ser"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论ID
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//文章是否存在
	_claims, _ := c.Get("claims")

	claims := _claims.(*jwts.CustomClaims)

	//文章是否存在
	_, err = es_ser.CommDetail(cr.ArticleID)

	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	//判断是否是字评论
	if cr.ParentCommentID != nil {
		//给父评论数 + 1
		var parentComment models.CommentModel
		// 找父评论
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			//发子评论，但是父评论不存在
			res.FailWithMessage("父评论不存在", c)
			return
		}
		// 判断父评论的文章是否和当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("评论文章不一致", c)
			return
		}
		// 给父评论的子评论数 + 1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + ?", 1))

	}
	// 拿到文章数，新的文字评论数存缓存里
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	// 给文章的评论数 + 1
	redis_ser.NewCommentCount().Set(cr.ArticleID)
	res.OKWithMessage("评论成功", c)

}
