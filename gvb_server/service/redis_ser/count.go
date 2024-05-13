package redis_ser

import (
	"gvb_server/gvb_server/global"
	"strconv"
)

type CountDB struct {
	Index string // 索引前缀
}

// Set设置某一个id的点赞数
func (c CountDB) Set(id string) error {
	number, _ := global.Redis.HGet(c.Index, id).Int()
	number = number + 1
	err := global.Redis.HSet(c.Index, id, number).Err()
	return err
}

// Set设置某一个id的值， 在原有的基础上加多少
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()
	return err
}

// Get获取某一个id的点赞数
func (c CountDB) Get(id string) int {
	number, _ := global.Redis.HGet(c.Index, id).Int()
	return number
}

// GetInfo获取所有点赞数
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

// Clear清空点赞数
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
