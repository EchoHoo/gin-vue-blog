package messageapi

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"`        //发送者ID
	SendUserNickName string    `json:"send_user_nick_name"` //
	SendUserAvatar   string    `json:"send_user_avatars"`   //
	RevUserID        uint      `json:"rev_user_id"`         //接收者ID
	RevUserNickName  string    `json:"rev_user_nick_name"`  //
	RevUserAvatar    string    `json:"rev_user_avatars"`    //
	Content          string    `json:"content"`             //消息内容
	CreatedAt        time.Time `json:"created_at"`          //最新的消息时间
	MessageCount     uint      `json:"message_count"`       //消息数量
}
type MessageGroup map[uint]*Message

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageGroup = MessageGroup{}

	var messageList []models.MessageModel

	var messages []Message

	err := global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id =? or rev_user_id = ?", claims.UserID, claims.UserID).Error

	if err != nil {
		res.FailWithMessage("获取消息列表失败", c)
		global.Log.Error(err)
		return
	}

	for _, model := range messageList {
		//判断是否为一个组，只要ID合相同就行
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevNickName,
			RevUserAvatar:    model.RevAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreateAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)

}
