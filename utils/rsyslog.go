/**
* Created by GoLand
* @file log.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date   2022/4/25 17:27
* @desc   write_log.go
 */

package utils

import (
	"fmt"
	"net"
	"strings"
)

const (
	UDP = "udp"
	TCP = "tcp"
)

type Rsyslog struct {
	con     net.Conn
	address string
	netType string
}

func (r *Rsyslog) Init(host string, port int, network ...string) *Rsyslog {
	r.address = fmt.Sprintf("%s:%d", host, port)
	r.netType = "udp"
	if len(network) > 0 {
		r.netType = "tcp"
	}
	return r
}

// Close 关闭连接
func (r *Rsyslog) Close() error {
	return r.con.Close()
}

// 初始化连接
func (r *Rsyslog) connect() error {
	conn, err := net.Dial(r.netType, r.address)
	if err != nil {
		return err
	}
	r.con = conn
	return nil
}

// 写日志
func (r *Rsyslog) Write(log string) (int, error) {
	if r.netType == UDP {
		if err := r.connect(); err != nil {
			return 0, err
		}
		log = strings.Trim(log, " \r\n")
		length, err := r.con.Write([]byte(log + "\n"))
		if err != nil {
			return 0, err
		}
		return length, err
	} else {
		if r.con == nil {
			if err := r.connect(); err != nil {
				return 0, err
			}
		}
		log = strings.Trim(log, " \r\n")
		length, err := r.con.Write([]byte(log + "\n"))
		if err != nil {
			return 0, err
		}
		return length, err
	}
}
