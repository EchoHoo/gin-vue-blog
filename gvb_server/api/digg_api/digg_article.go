package digg_api

import (
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//对ID的长度校验 或者 直接查ES是否能够直接找到
	redis_ser.NewDigg().Set(cr.ID)

	//redis_ser.Digg(cr.ID)

	res.OKWithMessage("文章点赞成功", c)
}
