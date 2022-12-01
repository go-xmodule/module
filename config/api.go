/**
 * Created by PhpStorm.
 * @file   ap.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:34
 * @desc   ap.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"log"
)

// Api 接口配置
type Api struct {
	Secret   string `yaml:"secret"`
	Overtime int64  `yaml:"overtime"`
}

func InitApiConfig(client config_client.IConfigClient, group string, config interface{}) {
	getConfigParams := nacos.GetConfigParams{
		Client: client,
		DataId: global.ApiConfigDataId,
		Group:  group,
	}
	err := nacos.GetConfig(getConfigParams, config)
	if err != nil {
		log.Fatal(global.GetApiConfigErr)
	}
}
