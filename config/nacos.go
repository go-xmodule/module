/**
 * Created by goland.
 * @file   nacos.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 21:41
 * @desc   nacos.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"os"
)

// NacosConfig Nacos 配置
type NacosConfig struct {
	Params Params `yaml:"params"`
}

type Params struct {
	Port        uint64 `yaml:"port"`
	NamespaceId string `yaml:"namespaceId"`
	LogDir      string `yaml:"logDir"`
	CacheDir    string `yaml:"cacheDir"`
	LogLevel    string `yaml:"logLevel"`
	Host        string `yaml:"host"`
}

// NacosConfigFile Nacos配置文件
const NacosConfigFile = "nacos.yaml"

// GetNacosConfig 获取系统配置
func GetNacosConfig() (NacosConfig, config_client.IConfigClient) {
	var nacosConfig NacosConfig
	err := utils.ParseConfig(NacosConfigFile, &nacosConfig)
	if utils.HasErr(err, global.GetSystemConfigErr) {
		os.Exit(403)
	}
	connectConfig := nacos.ConnectConfig{
		Host:        nacosConfig.Params.Host,
		Port:        nacosConfig.Params.Port,
		NamespaceId: nacosConfig.Params.NamespaceId,
		LogDir:      nacosConfig.Params.LogDir,
		CacheDir:    nacosConfig.Params.CacheDir,
		LogLevel:    nacosConfig.Params.LogLevel,
	}
	client, err := nacos.GetConfigClient(connectConfig)
	if err != nil {
		utils.Logger.Fatal(global.SystemInitFail.String())
	}
	return nacosConfig, client
}
