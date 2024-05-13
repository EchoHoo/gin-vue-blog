package ip_test

import (
	"fmt"
	"net"
	"testing"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func TestIp(t *testing.T) {
	fmt.Println(GetAddr("https://video.fengfengzhidao.com/play/1/147"))
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, err := db.City(parseIP)
	if err != nil {
		return "错误的IP地址"
	}
	if IsIntranetIP(parseIP) {
		return "内网IP地址"
	}
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}

	city := record.City.Names["zh-CN"]

	return fmt.Sprintf("%s-%s", province, city)
}

func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 == nil { // 错误的IP地址
		return true
	}

	//192.168
	//172.16 - 172.31
	//10
	//169.254
	return ip4[0] == 192 && ip4[1] == 168 || ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 || ip4[0] == 10 || ip4[0] == 169 && ip4[1] == 254
}
