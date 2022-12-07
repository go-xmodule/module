/**
 * Created by PhpStorm.
 * @file   mysql.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:06
 * @desc   mysql.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"log"
)

type Database struct {
	Type        string `yaml:"type"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	Database    Connect
}
type Connect struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"database"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timeZone"`
	Charset  string `yaml:"charset"`
	Mode     string `yaml:"mode"`
}

// Redis 配置
type Redis struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Db         int    `yaml:"db"`
	Password   string `yaml:"password"`
	MaxRetries int    `yaml:"maxRetries"`
}

func InitDatabaseConfig(client config_client.IConfigClient, group string, config interface{}) {
	getConfigParams := nacos.GetConfigParams{
		Client: client,
		DataId: global.DatabaseConfigDataId,
		Group:  group,
	}
	err := nacos.GetConfig(getConfigParams, config)
	if err != nil {
		log.Printf("%s,err:%s", global.GetConfigErr.String(), err.Error())
		log.Fatal(global.GetDbConfigErr)
	}
}
