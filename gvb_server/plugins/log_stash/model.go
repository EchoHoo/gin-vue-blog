package log_stash

import "time"

type LogStashModel struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	CreateAt time.Time `json:"created_at"`
	IP       string    `gorm:"size:32" json:"ip"`
	Addr     string    `gorm:"size:64" json:"addr"`
	Level    Level     `gorm:"size:4" json:"level"`     // 日志级别
	Content  string    `gorm:"size:128" json:"content"` // 日志内容
	UserID   uint      `json:"user_id"`                 // 登录用户的用户ID,需要自己查询的时候做关联
}
