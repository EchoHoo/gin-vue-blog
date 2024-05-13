package core

import (
	"gvb_server/gvb_server/global"
	"log"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatal("ip地址数据库加载失败", err)
	}
	global.AddrDB = db

}
