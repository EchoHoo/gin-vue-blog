package main

import (
	"gvb_server/gvb_server/core"
	_ "gvb_server/gvb_server/docs"
	"gvb_server/gvb_server/flag"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/routers"
	"gvb_server/gvb_server/service/cron_ser"
	"gvb_server/gvb_server/utils"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//初始化IP服务器
	core.InitAddrDB()
	defer global.AddrDB.Close()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//连接redis
	global.Redis = core.ConnectRedis()

	//连接ES
	global.ESClient = core.EsConnect()

	//初始化路由组
	router := routers.InitRouter()

	addr := global.Config.System.Addr()

	//打印系统信息
	utils.PrintSystem()

	//启动定时任务
	go cron_ser.CronInit()

	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalln(err.Error())
	}
}
