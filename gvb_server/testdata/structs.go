package testdata

import (
	"gvb_server/gvb_server/models"
)

type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"请输入广告标题" structs:"title"`
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法"`
	Images       string `json:"images" binding:"required,url" msg:"广告图片非法"`
	IsShow       bool   `json:"is_show" binding:"required" msg:"请选择是否展示"`
}

// func main() {
// 	u1 := AdvertRequest{
// 		Title:  "测试广告1",
// 		Href:   "http://www.baidu.com",
// 		Images: "http://www.baidu.com/img.jpg",
// 		IsShow: true,
// 	}
// 	m3 := structs.Map(&u1)
// 	fmt.Println(m3)

// }
