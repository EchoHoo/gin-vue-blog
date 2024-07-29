package images_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImagesApi) ImageNameListView(c *gin.Context) {

	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id,path,name").Scan(&imageList)

	res.OkWithData(imageList, c)
}
