package advert_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入广告标题" structs:"title"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`
	Images string `json:"images" binding:"required,url" msg:"广告图片非法" structs:"images"`
	IsShow bool   `json:"is_show" structs:"is_show"`
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @param data body AdvertRequest true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response "成功"
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复的判断，图片去重
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("广告标题重复", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OKWithMessage("添加广告成功", c)

}
