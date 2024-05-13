package main

import (
	"fmt"
	"gvb_server/gvb_server/core"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/service/redis_ser"
	"testing"
)

func TestDigg(t *testing.T) {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()

	digg := redis_ser.NewDigg()
	digg.Set("ws1c1Y4Bc16GF50VfZ6E")
	fmt.Println(digg.Get("ws1c1Y4Bc16GF50VfZ6E"))
	fmt.Println(digg.GetInfo())
}
