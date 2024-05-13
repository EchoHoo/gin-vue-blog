package tagapi

import (
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//判断Referer是否包含admin，如果是就全部返回，不是就返回is_show=true的广告
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	//需要展示该标签下的文章数量
	res.OkWithList(list, count, c)
}
