package menu_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	//先清空之间banner
	var menuModel models.MenuModel
	err = global.DB.Preload("Banners").Take(&menuModel, id).Error
	if err != nil {
		res.FailWithError(err, nil, c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()

	//再更新菜单,如果选择了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		//操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithError(err, nil, c)
			return
		}
	}
	//更新菜单
	//更新一些字段
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OKWithMessage("修改广告成功", c)
}
