package es_ser

import (
	"gvb_server/gvb_server/models"
)

type Option struct {
	models.PageInfo
	Fields []string
	Tag    string `form:"tag"`
}

func (o *Option) GetFrom() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
