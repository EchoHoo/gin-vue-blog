package log_stash

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/utils"
	"gvb_server/gvb_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func New(ip string, token string) *Log {
	//解析token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}
	addr := utils.GetAddr(ip)
	return &Log{
		ip:     ip,
		addr:   addr,
		userId: userID,
	}
}
func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)

}

type Log struct {
	ip     string
	addr   string
	userId uint
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

// 入库的一个方法
func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:       l.ip,
		Addr:     l.addr,
		UserID:   l.userId,
		Level:    level,
		Content:  content,
		CreateAt: time.Now(),
	}).Error
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(l.ip, l.userId, content, level)

}
