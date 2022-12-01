/**
 * Created by goland.
 * @file   link.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 09:58
 * @desc   link.go
 */

package nacos

import (
	"github.com/go-utils-module/module/global"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

type ConnectConfig struct {
	Host        string `json:"host,omitempty"`
	Port        uint64 `json:"port,omitempty"`
	NamespaceId string `json:"namespace_id,omitempty"`
	LogDir      string `json:"log_dir,omitempty"`
	CacheDir    string `json:"cache_dir,omitempty"`
	LogLevel    string `json:"log_level,omitempty"`
}

// GetNamingClient 服务连接配置
func GetNamingClient(config ConnectConfig) (naming_client.INamingClient, error) {
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig(config.Host, config.Port, constant.WithContextPath("/nacos")),
	}
	// create ClientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(config.NamespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(config.LogDir),
		constant.WithCacheDir(config.CacheDir),
		constant.WithLogLevel(config.LogLevel),
	)
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		log.Printf("%s,err:%s", global.GetNamingClientErr.String(), err.Error())
	}
	return client, err
}

// GetConfigClient 配置连接配置
func GetConfigClient(config ConnectConfig) (config_client.IConfigClient, error) {
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig(config.Host, config.Port, constant.WithContextPath("/nacos")),
	}
	// create ClientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(config.NamespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(config.LogDir),
		constant.WithCacheDir(config.CacheDir),
		constant.WithLogLevel(config.LogLevel),
	)
	// 创建动态配置客户端
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Printf("%s,err:%s", global.GetConfigClientErr.String(), err.Error())
	}
	return client, err
}
