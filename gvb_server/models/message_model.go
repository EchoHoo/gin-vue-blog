package models

//MessageModel 记录消息
type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"` //发送人ID
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avater"`

	RevUserID    uint      `gorm:"primaryKey" json:"rev_user_id"` //接收人ID
	RevUserModel UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevNickName  string    `gorm:"size:42" json:"rev_user_nick_name"`
	RevAvatar    string    `json:"rev_user_avatar"`
	IsRead       bool      `gorm:"default:false" json:"is_read"` //接受方是否查看
	Content      string    `json:"content"`                      //消息内容
}
