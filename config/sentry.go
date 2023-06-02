package config

import (
	"github.com/x-module/module/global"
	utils "github.com/x-module/utils/utils/config"
	"github.com/x-module/utils/utils/xlog"
)

// SentryConfigFile 配置文件
const SentryConfigFile = "sentry.yaml"

type SentryConfig struct {
	Params Params `yaml:"params"`
}

type Params struct {
	Environment      string  `yaml:"Environment"`
	ServerName       string  `yaml:"ServerName"`
	Release          string  `yaml:"Release"`
	Dsn              string  `yaml:"Dsn"`
	TracesSampleRate float64 `yaml:"TracesSampleRate"`
	AttachStacktrace bool    `yaml:"AttachStacktrace"`
	Debug            bool    `yaml:"Debug"`
}

// InitSentryConfig 获取Sentry配置
func InitSentryConfig() SentryConfig {
	var config SentryConfig
	path := utils.GetConfigFile(SentryConfigFile)
	err := utils.GetConfig(path, &config)
	if err != nil {
		xlog.Logger.Fatal(err, global.GetSentryConfigErr.String())
	}
	return config
}
