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

type RegisterServerParams struct {
	config       config.NacosConfig
	serverConfig config.Server
	serverName   string
	clusterName  string
	groupName    string
	Port         uint64
}

func RegisterServer(params RegisterServerParams) {
	connectConfig := center2.ConnectConfig{
		Host:        params.config.Params.Host,
		Port:        params.config.Params.Port,
		NamespaceId: params.config.Params.NamespaceId,
		LogDir:      params.config.Params.LogDir,
		CacheDir:    params.config.Params.CacheDir,
		LogLevel:    params.config.Params.LogLevel,
	}
	client, err := center2.GetNamingClient(connectConfig)
	if err != nil {
		utils.Logger.WithField(global.ErrField, err).Fatalln(global.GetNamingClientErr.String())
		return
	}
	registerServerParams := center2.RegisterServerParams{
		Client:      client,
		Port:        params.Port,
		ServiceName: params.serverName,
		ClusterName: params.clusterName,
		GroupName:   params.groupName,
		Metadata:    map[string]string{},
	}
	result, _ := center2.RegisterServer(registerServerParams)
	if result {
		utils.Logger.Debug("服务注册成功!")
	} else {
		utils.Logger.Fatalln(global.RegisterServerErr.String())
	}
}
