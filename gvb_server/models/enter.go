package models

type Page struct {
	Page  int    `form:"page"`  //当前页数
	Key   int    `form:"key"`   //关键字搜索
	Limit int    `form:"limit"` //获取多少页数
	Sort  string `form:"sort"`  //排序
}
type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list" form:"id_list" uri:"id_list"`
}

// 批量删除
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
