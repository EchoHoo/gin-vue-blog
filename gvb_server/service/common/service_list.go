package common

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug   bool
	Likes   []string // 模糊匹配的字段
	Preload []string // 预加载的字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "create_at desc" //默认按照时间
	}

	DB = DB.Model(model)
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	//这里的query会受到上面查询的影响，需要手动复位一下
	count = DB.Find(&list).RowsAffected

	query := DB.Where(model)
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	offset := (option.PageInfo.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit <= 0 {
		option.Limit = -1
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
