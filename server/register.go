/**
 * Created by goland.
 * @file   register.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 00:37
 * @desc   register.go
 */

package server

import (
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils"
	center2 "github.com/go-utils-module/module/utils/nacos"
)

type ConnectConfig struct {
	Host string `json:"host,omitempty"`
	Port uint64 `json:"port,omitempty"`
}

func RegisterServer(nacosConfig config.NacosConfig, serverConfig config.Server, serverName, clusterName string) {
	connectConfig := center2.ConnectConfig{
		Host:        nacosConfig.Params.Host,
		Port:        nacosConfig.Params.Port,
		NamespaceId: nacosConfig.Params.NamespaceId,
		LogDir:      nacosConfig.Params.LogDir,
		CacheDir:    nacosConfig.Params.CacheDir,
		LogLevel:    nacosConfig.Params.LogLevel,
	}
	client, err := center2.GetNamingClient(connectConfig)
	if err != nil {
		utils.Logger.WithField(global.ErrField, err).Fatalln(global.GetNamingClientErr.String())
		return
	}
	registerServerParams := center2.RegisterServerParams{
		Client:      client,
		Port:        uint64(serverConfig.Port),
		ServiceName: serverName,
		ClusterName: clusterName,
		Metadata:    map[string]string{},
	}
	result, _ := center2.RegisterServer(registerServerParams)
	if result {
		utils.Logger.Debug("服务注册成功!")
	} else {
		utils.Logger.Fatalln(global.RegisterServerErr.String())
	}
}
