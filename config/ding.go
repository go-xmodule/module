/**
 * Created by PhpStorm.
 * @file   ap.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:34
 * @desc   ap.go
 */

package config

import (
	"github.com/x-module/module/global"
	utils "github.com/x-module/utils/utils/config"
	"log"
)

const DingConfigFile = "dingding.yaml"

type DingConfig struct {
	Group Group `yaml:"group"`
}

type Group struct {
	Alarm  Alarm  `yaml:"alarm"`
	Notice Notice `yaml:"notice"`
}

type Alarm struct {
	Secret     string `yaml:"secret"`
	OpenApiUrl string `yaml:"openApiUrl"`
}

type Notice struct {
	OpenApiUrl string `yaml:"openApiUrl"`
	Secret     string `yaml:"secret"`
}

// InitDingConfig 钉钉配置
func InitDingConfig() DingConfig {
	var config DingConfig
	path := utils.GetConfigFile(DingConfigFile)
	err := utils.GetConfig(path, &config)
	if err != nil {
		log.Fatal(err, global.GetApiConfigErr.String())
	}
	return config
}
