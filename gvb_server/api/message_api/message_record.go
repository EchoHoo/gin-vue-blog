package messageapi

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户ID"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)

	err = global.DB.Order("create_at asc").
		Find(&_messageList, "send_user_id =? or rev_user_id = ?", claims.UserID, claims.UserID).Error

	if err != nil {
		res.FailWithMessage("获取消息列表失败", c)
		global.Log.Error(err)
		return
	}

	for _, model := range _messageList {
		//判断是否为一个组，只要ID合相同就行
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}

	//点开消息，里面的每一条消息都变成已读状态
	for _, model := range messageList {
		if model.RevUserID == claims.UserID {
			global.DB.Model(&models.MessageModel{}).Where("id = ?", model.ID).Update("is_read", 1)
		}
	}
	res.OkWithData(messageList, c)
}
