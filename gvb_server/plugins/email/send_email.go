package email

import (
	"gvb_server/gvb_server/global"

	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "验证码"
	Node  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}
func NewNode() Api {
	return Api{
		Subject: Node,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func send(name, subject, body string) error {

	e := global.Config.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}
func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) //发送者
	m.SetHeader("To", mailTo)                                //接收者
	m.SetHeader("Subject", subject)                          //主题
	m.SetBody("text/html", body)                             //正文
	d := gomail.NewDialer(host, port, userName, authCode)    //连接服务器
	return d.DialAndSend(m)
}
