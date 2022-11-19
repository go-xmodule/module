/**
 * Created by PhpStorm.
 * @file   server.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:37
 * @desc   server.go
 */

package config

// Server 服务配置
type Server struct {
	Protocol string `yaml:"protocol"`
	Domain   string `yaml:"domain"`
	Port     int    `yaml:"port"`
	Ip       string `yaml:"ip"`
}
