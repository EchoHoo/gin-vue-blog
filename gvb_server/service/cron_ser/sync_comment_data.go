package cron_ser

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/service/redis_ser"

	"gorm.io/gorm"
)

func SyncCommentData() {
	commentDiggInfo := redis_ser.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, "id = ?", key).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%d 更新成功 新的点赞数为 %d", comment.ID, comment.DiggCount+1)
	}
	redis_ser.NewCommentDigg().Clear()
}
