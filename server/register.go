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
	Config       config.NacosConfig `json:"config"`
	ServerConfig config.Server      `json:"server_config"`
	ServerName   string             `json:"server_name,omitempty"`
	ClusterName  string             `json:"cluster_name,omitempty"`
	GroupName    string             `json:"group_name,omitempty"`
	Port         uint64             `json:"port,omitempty"`
	Metadata     map[string]string  `json:"metadata"`
}

func RegisterServer(params RegisterServerParams) {
	connectConfig := center2.ConnectConfig{
		Host:        params.Config.Params.Host,
		Port:        params.Config.Params.Port,
		NamespaceId: params.Config.Params.NamespaceId,
		LogDir:      params.Config.Params.LogDir,
		CacheDir:    params.Config.Params.CacheDir,
		LogLevel:    params.Config.Params.LogLevel,
	}
	client, err := center2.GetNamingClient(connectConfig)
	if err != nil {
		utils.Logger.WithField(global.ErrField, err).Fatalln(global.GetNamingClientErr.String())
		return
	}
	registerServerParams := center2.RegisterServerParams{
		Client:      client,
		Port:        params.Port,
		ServiceName: params.ServerName,
		ClusterName: params.ClusterName,
		GroupName:   params.GroupName,
		Metadata:    params.Metadata,
	}
	result, _ := center2.RegisterServer(registerServerParams)
	if result {
		utils.Logger.Debug("服务注册成功!")
	} else {
		utils.Logger.Fatalln(global.RegisterServerErr.String())
	}
}
