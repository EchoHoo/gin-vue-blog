package log_api

import (
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/plugins/log_stash"
	"gvb_server/gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `json:"level" form:"level"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
}
