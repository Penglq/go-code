package main

import (
	"net"
	"fmt"
)

func main() {
	ip, mac := GetIPMac()
	fmt.Printf("IP:%s;MAC:%s", ip, mac)

}

func GetIPMac() (ip,mac string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for _, inter := range interfaces {
		x, _ := inter.Addrs()
		for _, a := range x {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					mac = inter.HardwareAddr.String()
					return
				}
			}
		}
	}
	return
}