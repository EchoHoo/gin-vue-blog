package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 //QQ
	SignGitee SignStatus = 2 //gitee
	SignEmail SignStatus = 3 //邮箱
)

func (s SignStatus) MashalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	default:
		str = "其他"
	}
	return str
}
