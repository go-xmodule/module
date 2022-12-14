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
	"fmt"
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

// OpenFreeUDPPort opens free UDP port
// This example does not actually use UDP ports,
// but to avoid port collisions with the HTTP server,
// it binds the same number of UDP port in advance.
func OpenFreeUDPPort(portBase int, num int) (net.PacketConn, int, error) {
	for i := 0; i < num; i++ {
		port := portBase + i
		conn, err := net.ListenPacket("udp", fmt.Sprint(":", port))
		if err != nil {
			continue
		}
		return conn, port, nil
	}
	return nil, 0, errors.New("failed to open free port")
}
