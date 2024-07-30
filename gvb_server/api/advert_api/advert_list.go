package advert_api

import (
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @param data query models.PageInfo     false "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	referer := c.GetHeader("Gvb_referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		//admin来的
		isShow = false
	}

	//判断Referer是否包含admin，如果是就全部返回，不是就返回is_show=true的广告
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
