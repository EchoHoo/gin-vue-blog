package jsonfilter

type User struct {
	Uit uint `json:uid,select(article)"`
}
