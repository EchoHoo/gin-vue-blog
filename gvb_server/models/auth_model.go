package models

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID       uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"` //主键ID
	CreateAt time.Time `json:"created_at,select($any)" structs:"-"`           //创建时间
	Update   time.Time `json:"-" structs:"-"`                                 //更新时间
}

// PageInfo paging common parameter structure
type PageInfo struct {
	Page  int      `json:"page" form:"page"`   //页码
	Limit int      `json:"limit" form:"limit"` //每页条数
	Key   string   `json:"key" form:"key"`     //关键字
	Sort  string   `json:"sort" form:"sort"`   //排序
	Like  []string `json:"like" form:"like"`   //需要模糊匹配的列表
}

// ResponseList 公共的返回
type ResponseList struct {
	List  any `json:"list"`
	Count any `json:"count"`
}

func (m *MODEL) BeforeSave(db *gorm.DB) error {
	if m.ID != 0 {
		m.Update = time.Now()
	} else {
		m.CreateAt = time.Now()
	}
	return nil
}
func (m *MODEL) BeforeUpdate(db *gorm.DB) error {
	if m.ID != 0 {
		m.Update = time.Now()
	} else {
		m.CreateAt = time.Now()
	}
	return nil
}
