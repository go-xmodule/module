/**
 * Created by goland.
 * @file   connect.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/7 15:34
 * @desc   connect.go
 */

package server

import (
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"log"
)

// GetNacosClient 获取系统配置
func GetNacosClient(nacosConfig config.NacosConfig) naming_client.INamingClient {
	connectConfig := nacos.ConnectConfig{
		Host:        nacosConfig.Params.Host,
		Port:        nacosConfig.Params.Port,
		NamespaceId: nacosConfig.Params.NamespaceId,
		LogDir:      nacosConfig.Params.LogDir,
		CacheDir:    nacosConfig.Params.CacheDir,
		LogLevel:    nacosConfig.Params.LogLevel,
	}
	client, err := nacos.GetNamingClient(connectConfig)
	if err != nil {
		log.Fatal(global.SystemInitFail.String())
	}
	return client
}
