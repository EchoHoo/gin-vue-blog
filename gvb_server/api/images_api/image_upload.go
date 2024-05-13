package images_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service"
	"gvb_server/gvb_server/service/image_ser"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

// 图片上传白名单
var (
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

// 上传图片,返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//multi是许多的意思，这里就是返回多个
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"] //对应的是form-data中的KEY,返回的是图片链表
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	//判读路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	//不存在就创建目录
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			// 本地还要保存一下
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}

	res.OkWithData(resList, c)
}
