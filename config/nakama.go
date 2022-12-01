/**
 * Created by PhpStorm.
 * @file   nakama.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 23:27
 * @desc   nakama.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
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

func InitNakamaConfig(client config_client.IConfigClient, group string) NakamaConfig {
	var nakamaConfig NakamaConfig
	getConfigParams := nacos.GetConfigParams{
		Client: client,
		DataId: global.NakamaConfigDataId,
		Group:  group,
	}
	err := nacos.GetConfig(getConfigParams, &nakamaConfig)
	if err != nil {
		log.Fatal(global.GetNakamaConfigErr)
	}
	return nakamaConfig
}
