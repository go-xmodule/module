/**
 * Created by goland.
 * @file   channel.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 20:45
 * @desc   channel.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	utils "github.com/go-utils-module/utils/utils/config"
	"log"
)

// ChannelConfigFile Channel配置文件
const ChannelConfigFile = "channel.yaml"

// InitChannelConfig Game统配置
func InitChannelConfig(config any) {
	path := utils.GetConfigFile(ChannelConfigFile)
	err := utils.GetConfig(path, config)
	if err != nil {
		log.Fatal(err, global.GetChannelConfigErr.String())
	}
}
