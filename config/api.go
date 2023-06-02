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

// Api 接口配置
type Api struct {
	Secret   string `yaml:"secret"`
	Overtime int64  `yaml:"overtime"`
}

// ApiConfigFile Api配置文件
const ApiConfigFile = "api.yaml"

// InitApiConfig Game统配置
func InitApiConfig(config any) {
	path := utils.GetConfigFile(ApiConfigFile)
	err := utils.GetConfig(path, config)
	if err != nil {
		log.Fatal(err, global.GetApiConfigErr.String())
	}
}
