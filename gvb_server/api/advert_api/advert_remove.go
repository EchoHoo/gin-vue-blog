package advert_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @param data body models.RemoveRequest true "批量删除广告"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object}  res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("没有找到相关广告", c)
		return
	}
	//批量删除
	global.DB.Delete(&advertList)
	res.OKWithMessage(fmt.Sprintf("共删除%d个广告", count), c)
}
