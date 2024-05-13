package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` //图片上传的大小 M
	Path string `yaml:"path" json:"path"` //图片上端的目录
}
