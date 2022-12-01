/**
 * Created by goland.
 * @file   game.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 20:46
 * @desc   game.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"log"
)

func InitGameConfig(client config_client.IConfigClient, group string, config interface{}) {
	getConfigParams := nacos.GetConfigParams{
		Client: client,
		DataId: global.GameConfigDataId,
		Group:  group,
	}
	err := nacos.GetConfig(getConfigParams, config)
	if err != nil {
		log.Fatal(global.GetGameConfigErr)
	}
}
