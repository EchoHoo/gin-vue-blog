package comment

import (
	"fmt"
	"gvb_server/gvb_server/core"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"testing"
)

func TestComment(t *testing.T) {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	fmt.Println(FindArticleCommentList("ws1c1Y4Bc16GF50VfZ6E"))

}
func FindArticleCommentList(articleID string) (RootCommentList []models.CommentModel) {

	global.DB.Find(&RootCommentList, "article_id =? and parent_comment_id is null", articleID)
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(model, &subCommentList)
	}
	return RootCommentList
}

// 递归查看子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}
