package menu_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type MenuDetailRequest struct {
	Path string `json:"path" form:"path"`
}

func (MenuApi) MenuDetailByPathView(c *gin.Context) {

	var cr MenuDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, "path = ?", cr.Path).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", menuModel.ID)
	banners := []Banner{}
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path})
	}
	menusResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menusResponse, c)
}
