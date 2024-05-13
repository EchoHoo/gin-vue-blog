package messageapi

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required" msg:"发送方不能为空"`
	RevUserID  uint   `json:"rev_user_id" binding:"required" msg:"接收方不能为空"`
	Content    string `json:"content" binding:"required" msg:"消息内容不能为空"`
}

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	//当前用户发布消息
	//sendid就是当前用户的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var sendUser models.UserModel
	var recvUser models.UserModel

	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送者不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收者不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserAvatar:   sendUser.Avatar,
		SendUserNickName: sendUser.NickName,

		RevUserID:   cr.RevUserID,
		RevAvatar:   recvUser.Avatar,
		RevNickName: recvUser.NickName,

		IsRead:  false,
		Content: cr.Content,
	}).Error
	if err != nil {
		res.FailWithMessage("发布消息失败", c)
		return
	}
	res.OKWithMessage("发布消息成功", c)
}
