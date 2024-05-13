package utils

import (
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

func GetIPList() (ipList []string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range interfaces {
		address, err := i2.Addrs()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, addr := range address {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ipNet.IP.To4() == nil {
				continue
			}
			ipList = append(ipList, ipNet.IP.String())
		}
	}
	return ipList
}
