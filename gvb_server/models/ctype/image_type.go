package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 //本地
	Qiniu ImageType = 2 //七牛云
)

func (s ImageType) MashalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}
