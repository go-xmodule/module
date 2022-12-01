/**
 * Created by goland.
 * @file   net.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 19:40
 * @desc   net.go
 */

package utils

import (
	"errors"
	"net"
)

// GetLocalIP 获取本机 IP
func GetLocalIP() (string, error) {
	ifAces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifAces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		address, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range address {
			ip := GetIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("unconnected to the network")
}

// GetIpFromAddr 获取ip
func GetIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
