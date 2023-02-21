/**
 * Created by PhpStorm.
 * @file   nakama.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 23:27
 * @desc   nakama.go
 */

package config

import (
	"github.com/go-xmodule/module/global"
	utils "github.com/go-xmodule/utils/utils/config"
	"log"
)

// NakamaConfig nakama 配置
type NakamaConfig struct {
	Account Account
	Base    Base
	Server  []string
}

// Account nakama账户信息
type Account struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	SignKey        string `yaml:"signKey"`
	ExpirationTime int64  `yaml:"expirationTime"`
}

// Base nakama 服务信息
type Base struct {
	ServerUrl string `yaml:"serverUrl"`
	Timeout   int    `yaml:"timeout"`
	RpcUrl    string `yaml:"rpcUrl"`
	Protocol  string `yaml:"protocol"`
	HttpKey   string `yaml:"httpKey"`
	SyncTimes int    `yaml:"syncTimes"`
	Port      int    `yaml:"port"`
}

// NakamaConfigFile Nakama配置文件
const NakamaConfigFile = "nakama.yaml"

// InitNakamaConfig 获取系统配置
func InitNakamaConfig() NakamaConfig {
	var nakamaConfig NakamaConfig
	path := utils.GetConfigFile(NakamaConfigFile)
	err := utils.GetConfig(path, &nakamaConfig)
	if err != nil {
		log.Fatal(err, global.GetNakamaConfigErr.String())
	}
	return nakamaConfig
}
