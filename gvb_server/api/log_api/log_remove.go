package log_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/plugins/log_stash"

	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var logList []log_stash.LogStashModel
	count := global.DB.Find(&logList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("没有找到相关日志", c)
		return
	}
	//批量删除
	global.DB.Delete(&logList)
	res.OKWithMessage(fmt.Sprintf("共删除%d个日志", count), c)
}
